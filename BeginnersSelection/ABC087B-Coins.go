package main

import (
	"fmt"
)

func main() {
	var a, b, c, x, cnt int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)

	if x%50 == 0 && (a+b+c) >= 1 && 0 <= a && a <= 50 && 0 <= b && b <= 50 && 0 <= c && c <= 50 && 50 <= x && x <= 20000 {
		for i := 0; i <= a; i++ {
			for j := 0; j <= b; j++ {
				for k := 0; k <= c; k++ {
					if i*500+j*100+k*50 == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
