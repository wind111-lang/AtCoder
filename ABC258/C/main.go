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

	var t,x int

	fmt.Fscan(r,&n,&q)
	fmt.Fscan(r,&s)

    p := 0
    str := strings.Split(s,"")

	for i := 0; i < q; i++{
		fmt.Fscan(r,&t,&x)
		if t == 1{
			p += x
		}else if t == 2{
            fmt.Fprintln(w,str[(x-p-1)%n])
		}
	}
}