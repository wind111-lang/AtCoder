package main

import(
	"fmt"
)

func main(){
	var a,b,c int
	fmt.Scanf("%d %d %d",&a,&b,&c)

	tmp := 0

	if a >= c{
		tmp = a
		a = c
		c = tmp
	}

	if(a <= b && b<=c){
		fmt.Println("Yes")
	}else{
		fmt.Println("No")
	}
}