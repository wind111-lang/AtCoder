package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	//var reel []int
	//sec := 0

	//flag := false

	fmt.Scanf("%d", &n)

	//cnt := [256]int{}

	s := make([]string, n)

	//reel := make([]int, n)

	// for i := 0; i < 9; i++ {
	reel := make([]int, 10)
	sum := make([]int, 10)
	// }

	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 10; j++ {
			pos := strings.Index(s[i], strconv.Itoa(j))
			fmt.Println(i, " in ", j, "pos: ", pos)
			reel[j] = pos
			sum[j] += reel[j]
		}
		fmt.Println("reel: ", reel)
		fmt.Println("reel sum: ", sum)
	}

	//fmt.Println(reel)
}
