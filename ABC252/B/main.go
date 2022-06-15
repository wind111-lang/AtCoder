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
	m := Atoi()

	max := 0
	cnt := 1

	flag := false

	a := make([]int, n)
	b := make([]int, m)

	for i := 0; i < n; i++ {
		a[i] = Atoi()
	}
	for i := 0; i < m; i++ {
		b[i] = Atoi()
		b[i] = b[i] - 1
	}

	for i := 0; i < n; i++ {
		if max < a[i] {
			max = a[i]
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//fmt.Println(a[j], b[i])
			//fmt.Println(len(a)-cnt, a[j], b[i])
			if len(a)-cnt == b[i] {
				if a[len(a)-cnt] == max {
					flag = true
					break
				}
			}
			//fmt.Println(cnt, len(a)-cnt)
			cnt++
		}

		cnt = 1

		if flag {
			break
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
