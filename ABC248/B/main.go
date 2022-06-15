package main

import (
	"fmt"
)

func main() {
	var a, b, k int
	cnt := 0

	fmt.Scan(&a, &b, &k)

	for a < b {

		a *= k
		cnt++
	}

	fmt.Println(cnt)
}
