/*=================================================================

Program name:
	selpg (SELect PaGes)[version-go]

Purpose:
	Sometimes one needs to extract only a specified range of
pages from an input text file. This program allows the user to do
that.

Author: Vasudev Ram
Transfer: ZhangZeMian

===================================================================*/

package main

/*================================= imports ======================*/
import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"io/ioutil"
	"io"
	"bufio"
	"flag"
)

/*================================= types =========================*/

type selpg_args struct
{
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type int
	print_dest string
}
var sp_args selpg_args

/*================================= globals =======================*/

var progname string
const int_max = 2147483647

/*================================= process_args() ================*/

func process_args(ac int, av []string, psa *selpg_args) {
	var s1 string
	var s2 string
	var argno int
	var i int

	if ac < 3 {
		fmt.Fprintf(os.Stderr,"%s: 1st arg should be -sstart_page\n", progname)
		usage()
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "DEBUG: before handling 1st arg\n")

	s1 = av[1]
	var s1_com = string(s1[0]) + string(s1[1])
	if strings.Compare(s1_com, "-s") != 0 {
		fmt.Fprintf(os.Stderr,"%s: 1st arg should be -sstart_page\n", progname)
		usage()
		os.Exit(2)
	}
	i,erri := strconv.Atoi(string(s1[2]));
	check(erri)
	if (i < 1) || (i > int_max - 1) {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %s\n", progname, string(s1[2]))
		usage()
		os.Exit(3)
	}
	psa.start_page = i

	s1 = av[2]
	s1_com = string(s1[0]) + string(s1[1])
	if strings.Compare(s1_com, "-e") != 0 {
		fmt.Fprintf(os.Stderr,"2nd arg should be -eend_page\n", progname)
		usage()
		os.Exit(4)
	}
	i,erri = strconv.Atoi(string(s1[2]));
	check(erri)
	if (i < 1) || (i > int_max - 1) {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %s\n", progname, string(s1[2]))
		usage()
		os.Exit(5)
	}
	psa.end_page = i

	fmt.Fprintf(os.Stderr, "DEBUG: before while loop for opt args\n");

	argno = 3
	var str string = av[argno]
	fmt.Fprintf(os.Stderr, "[str]:%s\n",str)
	for (argno <= (ac - 1) && rune(av[argno][0]) == '-'){
		s1 = av[argno]
		fmt.Fprintf(os.Stderr, "op=%c\n", rune(s1[1]));
		switch rune(s1[1]) {
			case 'l':
				rs := []rune(s1)
				s2 = string(rs[2:len(rs)])
				i, err_i := strconv.Atoi(s2)
				check(err_i)
				if (i < 1 || i > int_max - 1){
					fmt.Fprintf(os.Stderr,"%s: invalid page length %s\n", progname, s2)
					usage()
					os.Exit(6)
				}
				psa.page_len = i
				argno+=1
				continue
				break

			case 'f':
				if strings.Compare(s1, "-f") != 0{
					fmt.Fprintf(os.Stderr,"%s: option should be \"-f\"\n", progname)
					usage()
					os.Exit(7)
				}
				psa.page_type = 'f'
				argno+=1
				continue
				break

			case 'd':
				rs := []rune(s1)
				s2 = string(rs[2:len(s1) - 2])
				if (len(s2)<1){
					fmt.Printf("%s: -d option requires a printer destination\n", progname)
					usage()
					os.Exit(8)
				}
				psa.print_dest = s2
				argno+=1
				continue
				break

			default:
				fmt.Fprintf(os.Stderr, "%s: unknown option %s\n", progname, s1)
				usage()
				os.Exit(9)
				break
		}

	}

	fmt.Printf("DEBUG: before check for filename arg\n")
	fmt.Printf("DEBUG: argno = %d\n", argno)

	if argno <= (ac - 1){
		fmt.Fprintf(os.Stderr,"%s: filename\n", av[argno])
		psa.in_filename = av[argno]
		_, err_o := os.Stat(psa.in_filename)
		if os.IsNotExist(err_o){
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n",progname, psa.in_filename);
			os.Exit(10);
		}
		_, err_r := ioutil.ReadFile(psa.in_filename)
		if err_r != nil{
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" exists but cannot be read\n",
			progname, psa.in_filename);
			os.Exit(11);
		}
	}

	if !(psa.start_page > 0 &&
	(psa.end_page > 0 && psa.end_page >= psa.start_page) && 
	psa.page_len > 1 &&
	(psa.page_type == 'l' || psa.page_type == 'f')){
		os.Exit(12)
	}

	fmt.Printf("DEBUG: psa->start_page = %d\n", psa.start_page)
	fmt.Printf("DEBUG: psa->end_page = %d\n", psa.end_page)
	fmt.Printf("DEBUG: psa->page_len = %d\n", psa.page_len)
	fmt.Printf("DEBUG: psa->page_type = %c\n", psa.page_type)
	fmt.Printf("DEBUG: psa->print_dest = %s\n", psa.print_dest)
	fmt.Printf("DEBUG: psa->in_filename = %s\n", psa.in_filename)
}

/*================================= process_input() ===============*/

func check(e error) {
	if e != nil{
		panic(e)
	}	
}

func process_input(sa selpg_args) {
	var fin *os.File
	var fout *os.File
	var line string
	var line_ctr int
	var page_ctr int
	fmt.Fprintf(os.Stderr, "[open]:%s\n",sa.in_filename)

	fin, err := os.Open(sa.in_filename)
	fmt.Fprintf(os.Stderr, "[open]:%s\n",sa.in_filename)
	check(err)
	if len(sa.print_dest) == 0 {
		fout = os.Stdout;
	} else {
		_, err_f := os.Stat(sa.print_dest)
		if os.IsNotExist(err_f){
			fout, err_f = os.Create(sa.print_dest)
		} else{
			fout, err_f = os.OpenFile(sa.print_dest, os.O_APPEND,0666)
		}
		check(err_f)
	}
	//fmt.Sprintf(s1, "lp -d%s", sa.print_dest)
	//fout, err = os.Create(s1)
	//check(err)

	line_ctr = 0
	page_ctr = 1
	buf := bufio.NewReader(fin)
	for true {
		line,err = buf.ReadString('\n')
		if err == io.EOF{
			break
		}
		line_ctr++
		if line_ctr > sa.page_len{
			page_ctr++
			line_ctr = 1
		}
		if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page){
			fmt.Fprintf(fout, "%s", line)
		}
	}
	if page_ctr < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Fprintf(os.Stderr,"%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.end_page, page_ctr);
	}
	
}

/*================================= usage() =======================*/

func usage() {
	fmt.Fprintf(os.Stderr,
	"\nUSAGE: %s -s=start_page -e=end_page [ -f=true | -l=lines_per_page ][ -d=dest ] [ -i=in_filename ]\n", progname);
}


func main() {
	var s int
	flag.IntVar(&s, "s", -1, "Use -Int")
	var e int
	flag.IntVar(&e, "e", -1, "Use -Int")
	var f bool
	flag.BoolVar(&f, "f", false, "Use -Bool")
	var l int
	flag.IntVar(&l, "l", 72, "Use -Int")
	var d string
	flag.StringVar(&d, "d", "", "Use -String")
	var i string
	flag.StringVar(&i, "i", "", "Use -String")
	flag.Parse()
	var av []string
	for i := 0; i < len(os.Args); i++{
		str := string(os.Args[i])
		str = strings.Replace(str,"=","",-1)
		str = strings.Replace(str,"true","",-1)
		str = strings.Replace(str,"-i","",-1)
		av = append(av, str)
		fmt.Printf("%s\n", av[i])
	}

	ac := len(os.Args)

	var sa selpg_args


	progname = av[0]
	

	sa.start_page = -1
	sa.in_filename = ""
	sa.print_dest = ""
	sa.end_page = -1
	sa.page_len = 72
	sa.page_type = 'l'

	process_args(ac, av, &sa)
	fmt.Fprintf(os.Stderr, "[sucess1]:%d\n",1)
	process_input(sa)
	fmt.Fprintf(os.Stderr, "[sucess2]:%d\n",2)
}

