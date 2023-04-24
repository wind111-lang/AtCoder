package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

	var n,q int
	var l,num int
	var s,t int

	fmt.Fscan(r,&n,&q)
	arr := make([][]int, 0,n)

	for i:=0; i<n; i++{
		fmt.Fscan(r,&l)
		tmp := make([]int,0,l)
		for j:=0; j<l; j++{
			fmt.Fscan(r,&num)
			tmp = append(tmp, num)
		}
		arr = append(arr, tmp)
	}
	for i:=0; i<q; i++{
		fmt.Fscan(r,&s,&t)
		fmt.Fprintln(w,arr[s-1][t-1])
	}
}