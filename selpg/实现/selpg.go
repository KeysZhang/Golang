package main

/*================================= imports ======================*/
import (

	"fmt"
	"flag"
	"os"
	"io/ioutil"
	"io"
	"bufio"
	"os/exec"

)

/*================================= types =========================*/

type selpg_args struct
{
	start_page int
	end_page int
	in_filename string
	page_len int
	page_type bool
	print_dest string
}

/*================================= fun =========================*/

func main(){

	var args selpg_args
	get_args(&args)
	check_args(&args)
	run(&args)
	
}

func get_args(parg *selpg_args) {

	flag.IntVar(&(parg.start_page), "s", -1, "startPage")
	flag.IntVar(&(parg.end_page), "e", -1, "endPage")
	flag.IntVar(&(parg.page_len), "l", 72, "pageLength")
	flag.StringVar(&(parg.print_dest), "d", "", "printDest")
	flag.BoolVar(&(parg.page_type), "f", false, "pageType")

	flag.Parse()

	args_left := flag.Args()
	if(len(args_left) > 0){
		parg.in_filename = string(args_left[0])
	} else {
		parg.in_filename = ""
	}
}

func check_args(parg *selpg_args) {

	if parg == nil{
		fmt.Fprintf(os.Stderr, "\n[Error]The args is nil!Please check your program!\n\n")
		os.Exit(1)
	}else if(parg.start_page == -1) || (parg.end_page == -1){
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage is not allowed empty!Please check your command!\n\n")
		os.Exit(2)
	}else if (parg.start_page < 0) || (parg.end_page < 0){
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage and endPage is not negative!Please check your command!\n\n")
		os.Exit(3)
	}else if parg.start_page > parg.end_page{
		fmt.Fprintf(os.Stderr, "\n[Error]The startPage can not be bigger than the endPage!Please check your command!\n\n")
		os.Exit(4)
	}else{
		pt := 'f'
		if parg.page_type == false {
			pt = 'l'
		}
		fmt.Printf("\n[ArgsStart]\n")
		fmt.Printf("startPage: %d\nendPage: %d\ninputFile: %s\npageLength: %d\npageType: %c\nprintDestation: %s\n[ArgsEnd]\n\n", parg.start_page, parg.end_page, parg.in_filename, parg.page_len, pt, parg.print_dest)
	}

}

func run(parg *selpg_args) {
	
	var fin *os.File
	var fout *os.File
	var fout_d io.WriteCloser

	if parg.in_filename == ""{
		fin = os.Stdin
	} else {
		check_file_access(parg.in_filename)
		var err_fin error
		fin, err_fin = os.Open(parg.in_filename)
		check_err(err_fin, "fileInput")
	}

	fout, fout_d = check_fout(parg.print_dest)

	if(fout != nil){
		output_to_file(fout, fin, parg.start_page, parg.end_page, parg.page_len)
	} else {
		output_to_exc(fout_d, fin, parg.start_page, parg.end_page, parg.page_len)
	}

}

func check_file_access(filename string) {
	
	_, err_o := os.Stat(filename)
	if os.IsNotExist(err_o){
		fmt.Fprintf(os.Stderr, "\n[Error]: input file \"%s\" does not exist\n\n",filename);
		os.Exit(5);
	}
	_, err_r := ioutil.ReadFile(filename)
	check_err(err_r, filename)
}

func check_err(err error, object string) {
	if err != nil{
		fmt.Fprintf(os.Stderr, "\n[Error]%s:",object);
		panic(err)
	}
}

func check_fout(printDest string) (*os.File, io.WriteCloser) {

	var fout *os.File
	var fout_d io.WriteCloser

	if len(printDest) == 0 {
		fout = os.Stdout;
		fout_d = nil
	} else {
<<<<<<< HEAD
		_, err_f := os.Stat(sa.print_dest)
		if os.IsNotExist(err_f){
			fout, err_f = os.Create(sa.print_dest)
		} else{
			fout, err_f = os.OpenFile(sa.print_dest, os.O_APPEND,0666)
		}
		check(err_f)
	})
=======
		fout = nil
		var err_dest error
		cmd := exec.Command("./" + printDest)
		fout_d, err_dest = cmd.StdinPipe()
		check_err(err_dest, "fout_dest")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err_start := cmd.Start() 
		check_err(err_start,"command-start")
	}
>>>>>>> a780e752738b5cb79bc38dcc36e7e719acbde390

	return fout, fout_d
}

func output_to_file(fout *os.File, fin *os.File, pageStart int, pageEnd int, pageLen int) {

	line_ctr := 0
	page_ctr := 1
	buf := bufio.NewReader(fin)
	for true {
		line,err := buf.ReadString('\n')
		if err == io.EOF{
			break
		}
		check_err(err, "file_in_out")
		line_ctr++
		if line_ctr > pageLen{
			page_ctr++
			line_ctr = 1
		}
		if (page_ctr >= pageStart) && (page_ctr <= pageEnd){
			fmt.Fprintf(fout, "%s", line)
		}
	}
	if page_ctr < pageStart {
		fmt.Fprintf(os.Stderr, "\n[Error]: start_page (%d) greater than total pages (%d), no output written\n\n", pageStart, page_ctr)
		os.Exit(6)
	} else if page_ctr < pageEnd {
		fmt.Fprintf(os.Stderr,"\n[Error]: end_page (%d) greater than total pages (%d), less output than expected\n\n", pageEnd, page_ctr);
		os.Exit(7)
	}
}

func output_to_exc(fout io.WriteCloser, fin *os.File, pageStart int, pageEnd int, pageLen int) {

	line_ctr := 0
	page_ctr := 1
	buf := bufio.NewReader(fin)

	for true {
		bytes, err := buf.ReadByte()
		if err == io.EOF{
			break
		}
		if line_ctr > pageLen{
			page_ctr++
			line_ctr = 1
		}
		if (page_ctr >= pageStart) && (page_ctr <= pageEnd){
			fout.Write([]byte{bytes})
		}
	}
	if page_ctr < pageStart {
		fmt.Fprintf(os.Stderr, "\n[Error]: start_page (%d) greater than total pages (%d), no output written\n\n", pageStart, page_ctr)
		os.Exit(8)
	} else if page_ctr < pageEnd {
		fmt.Fprintf(os.Stderr,"\n[Error]: end_page (%d) greater than total pages (%d), less output than expected\n\n", pageEnd, page_ctr);
		os.Exit(9)
	}
}