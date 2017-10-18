package main

import(

	"bufio"
	"io"
	"os"
	"fmt"

)

func main() {
	var fout *os.File
	fin := os.Stdin
	buf := bufio.NewReader(fin)
	print_dest:="wf_output.txt"
	_, err_f := os.Stat(print_dest)
	if os.IsNotExist(err_f){
		fout, err_f = os.Create(print_dest)
	} else{
		fout, err_f = os.OpenFile(print_dest, os.O_APPEND,0666)
	}
	for true {
		line,err := buf.ReadByte()
		if err == io.EOF{
			break
		}
		fmt.Fprintf(fout, "%s", line)
	}
}