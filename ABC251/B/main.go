package main

import (
	"fmt"
)

func main() {
	var n, w int
	cnt := 0
	// sum := 0
	// flag := 0

	// restrict := 0

	fmt.Scanf("%d %d", &n, &w)

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])

		// 	if n == 2 {
		// 		if a[i] < w {
		// 			cnt++
		// 			restrict++
		// 		}
		// 		if flag == 0 {
		// 			if a[0]+a[1] < w {
		// 				cnt++
		// 				flag = 1
		// 			}
		// 		}
		// 	} else {
		// 		sum += a[i]
		// 		if restrict == 3 {
		// 			break
		// 		} else {
		// 			if a[i] < w {
		// 				cnt++
		// 				restrict++
		// 			}
		// 		}
		// 	}
	}

	fmt.Println(cnt)
}
