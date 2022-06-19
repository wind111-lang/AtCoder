package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64

	fmt.Scan(&n)
	pow := math.Pow(2, n)
	fmt.Println(int(pow))

	//fmt.Println(n)
}

//別解

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var n int
// 	pow := 1

// 	fmt.Scan(&n)
// 	for i := 0; i < n; i++ {
// 		pow = pow << 1
// 	}

// 	fmt.Println(pow)
// }