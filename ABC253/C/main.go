package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, q, x, c int
	var temp []int
	fmt.Scan(&n)

	s := make([]int, 0, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&q)
		if(q==1){
			fmt.Scan(&x)
			s = append(s, x)
			//fmt.Println(s)
		}else if(q==2){
			fmt.Scan(&x)
            fmt.Scan(&c)

			for c > 0{
				if s[0] == x {
					s2 := s[1:]
					s = s2
				}
				c--
			}
		}else{
			fmt.Println(s[0] - s[len(s)-1])
		}
		sort.Ints(s)
	}
}
