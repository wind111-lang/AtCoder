package main

import (
	"fmt"
)

func main() {
	var n,x int
	//var a,b,c,d int

	p := 0

	fmt.Scan(&n)

	slice := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		slice = append(slice, x)
	}

	for mass := range slice {
		fmt.Println(mass)
	}

	fmt.Println(p)
}