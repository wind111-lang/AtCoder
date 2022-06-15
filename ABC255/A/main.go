package main

import(
	"fmt"
)

func main()  {
	var r,c int
	var a1,a2 int
	matrix := make([][]int,0,2)

	cnt := 0
	num := 0
	fmt.Scan(&r,&c)

	for i:=0; i<2; i++{
		fmt.Scan(&a1,&a2)
		matrix = append(matrix, []int{a1,a2})
	}

	//fmt.Println(matrix)
	//fmt.Println(r+c)
	// for i:=0; i<2; i++{
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Println(matrix[i][j])
	// 	}
	// }

	for i:=0; i<2; i++{
		for j:=0; j<2; j++{
			if((i+1)==r && (j+1)==c){
				num = matrix[i][j]
				//fmt.Println(matrix[i][j])
				break
			}
			cnt++
		}
	}
	//fmt.Println(matrix)
	fmt.Println(num)
}