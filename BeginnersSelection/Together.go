package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
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
	var x int
	var count int
	var max int

	rand.Seed(time.Now().UnixNano())

	sc.Split(bufio.ScanWords)
	n := Atoi()

	a := make([]int, n, n)

	if n == len(a) {
		for i := 0; i < n; i++ {
			a[i] = Atoi()
			sel := rand.Intn(3)
			switch sel {
			case 0:
				a[i] += 1
			case 1:
				a[i] -= 1
			}
		}

		//b := a
		sort.Sort(sort.Reverse(sort.IntSlice(a)))

		for i := 1; i < n; i++ {
			if a[i-1] == a[i] {
				count++
			} else {
				if max < count {
					max = count
					x = a[i-1]
				}
				count = 0
			}
		}

		if x != 0 {
			fmt.Println(max)
		}
	}
}
