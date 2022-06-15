package main

import (
	"fmt"
)

func main() {
	var n string

	fmt.Scanf("%s", &n)
	tmp := n

	for len(n) < 6 {
		n += tmp
	}
	fmt.Println(n)
}
