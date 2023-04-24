package main

import (
	"fmt"
)

func main() {
	var N, D int
	var T []int

	fmt.Scan(&N)
	fmt.Scan(&D)

	for i := 0; i < N; i++ {
		var t int
		fmt.Scan(&t)
		T = append(T, t)
	}
	for i := 1; i < N; i++ {
		if (T[i] - T[i-1]) <= D {
			fmt.Print(T[i])
			return
		}
	}
	fmt.Print(-1)
}
