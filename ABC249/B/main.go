package main

import (
	"fmt"
	"unicode"
)

func main() {
	var S string
	var up, low int
	flag := 1

	fmt.Scan(&S)

	ui := [256]int{}

	for _, v := range S {
		//fmt.Println(ui[v])
		if unicode.IsUpper(v) {
			up++
		}
		if unicode.IsLower(v) {
			low++
		}
		if ui[v] != 0 {
			flag = 0
		}
		ui[v]++
	}

	if flag != 0 && up > 0 && low > 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

	//fmt.Println(up, low)
}
