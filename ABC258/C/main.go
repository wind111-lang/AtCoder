package main

import (
	"bufio"
	"fmt"
	"os"
    "strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

	var n,q int
	var s string

	//slice := make([]string, 0,n)

	var t,x int

	fmt.Fscan(r,&n,&q)
	fmt.Fscan(r,&s)

    str := strings.Split(s,"")

    num := 0

	//fmt.Println(bytes)

	//fmt.Println(s_arr)

	for i := 0; i < q; i++{
		fmt.Fscan(r,&t,&x)
		if t == 1{
          num += x
		}else if t == 2{
          fmt.Fprintln(w,string(str[(x-num-1)%n]))
		}
	}
}