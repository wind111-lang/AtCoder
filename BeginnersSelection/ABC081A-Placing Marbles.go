package main

import (
	"fmt"
	"strings"
)

func main() {
	var c int
	var sub string
	fmt.Scanf("%s", &sub)

	s := strings.Split(sub, "")

	if s[0] == "1" {
		c++
	}
	if s[1] == "1" {
		c++
	}
	if s[2] == "1" {
		c++
	}
	fmt.Println(c)
}
