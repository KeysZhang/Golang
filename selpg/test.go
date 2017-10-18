package main
import (
	"fmt"
	"strconv"
)
func main() {
	var s string = "1hello"
	//var a = string(s[0])
	b,err := strconv.Atoi(s[0])
	fmt.Printf("%d  %s", b, err)
}