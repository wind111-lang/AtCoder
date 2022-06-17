package main

import (
	"fmt"
	_"math"
)

func main() {
	var n, a, b int
	var x [][]string
	//var val string
	cnt := 0
	fmt.Scanf("%d %d %d", &n, &a, &b)

	for i := 0; i < n*a; i++ {
		for j := 0; j<n*b; j++{
			if ((i/a)+(j/b))%2 == 0 {
				x = append(x, []string{"."})
			} else {
				x = append(x, []string{"#"})
			}
		} 
	}
	for val := range x {
		if cnt==(n*b) {
			fmt.Println()
			cnt = 0
		}
		fmt.Print(x[val][0])
		cnt++
	}
	fmt.Println()
}
