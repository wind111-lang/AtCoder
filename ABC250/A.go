package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w, r, c int
	var cnt int

	fmt.Scanf("%d %d", &h, &w)
	fmt.Scanf("%d %d", &r, &c)

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			if (math.Abs(float64(i-r)))+(math.Abs(float64(j-c))) == 1 {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
