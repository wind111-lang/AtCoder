package main

import(
    "fmt"
)

func main(){
	var n int
	fmt.Scan(&n)

	pascal := make([]int, n)
	pascal[0] = 1

	for i := 0; i < n; i++ {
		for j := i; j > 0; j-- {
			pascal[j] += pascal[j-1]
		}

		for _,num := range pascal[0:i+1] {
			fmt.Print(num, " ")
		}
		fmt.Println()
	}
}