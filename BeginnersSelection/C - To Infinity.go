package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	var k int

	fmt.Scanf("%s", &s)
	fmt.Scanf("%d", &k)

	if k == 0 {
		fmt.Println("1")
	} else {
		arr := strings.Split(s, "")

		for i := 0; i < k; i++ {
			fmt.Println(arr[i])
		}
	}
}
