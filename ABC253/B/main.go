package main

import(
	"fmt"
)

func main(){
	var h int
	var w int

	mass := make([]string,0,h)

	fmt.Scanf("%d %d",&h,&w)

	for i:=0;i<h;i++{
		fmt.Scan(&mass[i])
		mass = append(mass,mass[i])
	}

}