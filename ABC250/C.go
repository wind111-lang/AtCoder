package main

import (
	"fmt"
	"bufio"
	"os"
)

func swap(a int, b int) (int, int) {
	return b, a
}

func main() {
	r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

	var n, q int

	fmt.Fscan(r, &n)

	val := make([]int, n+1)
	pos := make([]int, n+1)

	for i := 1; i <= n; i++ {
		val[i], pos[i] = i, i
		//fmt.Println(val[i], pos[i])
	}

	//fmt.Println(list, len(list))

	fmt.Fscan(r, &q)

	x := make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Fscan(r, &x[i])
	}

	for i := 0; i < q; i++ {
		p0 := pos[x[i]]
		p1 := p0

		if p0 != n {
			p1++
		} else {
			p1--
		}

		v0 := val[p0]
		v1 := val[p1]

		//fmt.Println(val[p0], val[p1])
		val[p0], val[p1] = swap(val[p0], val[p1])
		//fmt.Println(val[p0], val[p1])
		//fmt.Println(pos[v0], pos[v1])
		pos[v0], pos[v1] = swap(pos[v0], pos[v1])
		//fmt.Println(pos[v0], pos[v1])
	}

	for i := 1; i <= n; i++ {
		if i != 1 {
			fmt.Print("")
		}
		fmt.Fprintln(w,val[i])
	}
	//fmt.Println(val)

	// for i := 1; i <= q; i++ {
	// 	fmt.Scanf("%d", &x)
	// 	if x == n {
	// 		//arr[x-2], arr[x-1] = swap(arr[x-2], arr[x-1])
	// 	} else {
	// 		//arr[x-1], arr[x] = swap(arr[x-1], arr[x])
	// 	}
	// 	//fmt.Println(arr)
	// }
	// //fmt.Println(arr)
}
