package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func Atoi() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		log.Fatal(e)
	}
	return i
}

func main() {
	var res int

	sc.Split(bufio.ScanWords)
	n := Atoi()

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		a[i] = Atoi()
	}

	for i := 0; i < n; i++ {
		res += a[i] - 1
	}
	//fmt.Println(a)
	fmt.Println(res)
}
