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
	b := make([]int, n, n)

	var hit, blow int

	for i := 0; i < n; i++ {
		a[i] = Atoi()
	}
	for i := 0; i < n; i++ {
		b[i] = Atoi()
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] && i == j {
				hit++
				break
			} else if a[i] == b[j] {
				blow++
				break
			}
		}
	}
	fmt.Println(hit)
	fmt.Println(blow)
}
