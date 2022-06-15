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
	sc.Split(bufio.ScanWords)
	n := Atoi()

	a := make([]int, n, n)

	for i := 0; i < n; i++ {
		a[i] = Atoi()
	}

	flag := true

	res := 0

	for {
		for i := 0; i < n; i++ {
			if a[i]%2 != 0 {
				flag = false
			}
		}
		if flag != true {
			break
		}
		for i := 0; i < n; i++ {
			a[i] /= 2
		}
		res++
	}
	fmt.Println(res)
}
