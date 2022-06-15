package main

import (
	"fmt"
)

func main() {
	var n int
	var a int

	var list []int

	fmt.Scanf("%d", &n)

	set := make(map[int]bool)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		list = append(list, a)
	}

	for _, v := range list {
		if _, t := set[v]; t {
			set[v] = false
			delete(set, v)
		} else {
			set[v] = true
		}
	}
	fmt.Println(len(set))
}
