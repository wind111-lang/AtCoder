package main

import (
	"fmt"
	"math"
)

func main() {
	var n, a, b int

	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 1; i <= n*a; i++ {
		if (math.Abs(float64(i-a)))+(math.Abs(float64(i-b))) == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
		for j := 1; j < n*b; i++ {
			if (math.Abs(float64(j-a)))+(math.Abs(float64(j-b))) == 1 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
