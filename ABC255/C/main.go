package main

import(
	"fmt"
	"math"
)


func main()  {
	var x,a,d,n  int

	fmt.Scan(&x,&a,&d,&n)
	arr := make([]int, 0, n)
	//abs_arr := make([]Abs, 0,n)

	tmp := a
	abs := 0.0
	cnt := 0

	//arr = append(arr, tmp)

	for i := 0; i < n; i++{
		arr = append(arr,tmp)
		tmp += d
	}
	min_arr := arr[0]
	min_abs := math.Abs(float64(arr[0]) - float64(x) )

	for i := 0; i < len(arr); i++ {
		abs = math.Abs(float64(arr[i]) - float64(x) )
		//fmt.Println(abs)
		//abs_arr = append(abs_arr, Abs{abs,arr[i]})

		if min_abs > abs {
			min_abs = abs
			min_arr = arr[i]
		}
	}
	//fmt.Println(min_arr)
	if d != 0{
	    for{
		    if min_arr > x{
		    min_arr -= 1
		    }else{
			    min_arr += 1
		    }
		    cnt++
		    if x == min_arr{
			    break
		    }
	    }
    }
	fmt.Println(cnt)
}