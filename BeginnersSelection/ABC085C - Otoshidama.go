package main

import (
	"fmt"
)

func main() {

	var n, y int
	fmt.Scanf("%d %d", &n, &y)

	a := 0
	b := 0
	c := 0

	for i := 0; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			k := n - i - j
			if i*10000+j*5000+k*1000 == y {
				a = i
				b = j
				c = k
				fmt.Println(i, j, k)
				return
			}
		}
	}
	if a == 0 && b == 0 && c == 0 {
		fmt.Println("-1 -1 -1")
	}
}
