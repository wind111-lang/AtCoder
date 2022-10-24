package main

import(
	"fmt"
)

func main() {
	var a,b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Printf("%.3f\n",b/a)
}