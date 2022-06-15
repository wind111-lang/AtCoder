package main

import(
	"fmt"
	"math"
)

func main()  {
	var n, k int
	var x,y float64
	var a int
	//var r float64

	fmt.Scan(&n, &k)

	light := make([]int, 0, k)
	len_x := make([]float64, 0, n)
	len_y := make([]float64, 0, n)

	for i := 0; i < k; i++ {
		fmt.Scan(&a)
		a--
		light = append(light, a)
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&x, &y)
		len_x = append(len_x, x)
		len_y = append(len_y, y)
	}

	//fmt.Println(lenx, len_y)
	res := 0.0
	for i := 0; i < n; i++ {
		cres := math.MaxFloat64
		for nx := range light{
			cres = math.Min(cres,(len_x[i]-len_x[nx])*(len_x[i]-len_x[nx])+(len_y[i]-len_y[nx])*(len_y[i]-len_y[nx]))
		}
		res = math.Max(res,cres)
	}

    fmt.Printf("%.12f",math.Sqrt(res))

	//fmt.Println(light, pos)
}