package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var k int

	fmt.Scanf("%d", &k)
	str := strconv.FormatUint(uint64(k), 2)
	//fmt.Println(s)
	i := strings.Replace(str, "1", "2", -1)
	fmt.Println(i)
}
