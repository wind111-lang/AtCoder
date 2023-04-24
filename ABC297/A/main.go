package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	var S string

	fmt.Scan(&N)
	fmt.Scan(&S)

	if strings.Contains(S, "o") && !strings.Contains(S, "x") {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
