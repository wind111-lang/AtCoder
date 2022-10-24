package main

import(
	"fmt"
	"strings"
)

func main(){
	var n,m int
	var mass string
	count := 0

	fmt.Scan(&n)
	fmt.Scan(&m)

	graph := make([][]string,0,m)

	for i:=0; i<n; i++{
		fmt.Scan(&mass)
		graph = append(graph,strings.Split(mass,""))
	}

	for i:=0; i<m; i++{
		for j:=0; j<n; j++{
			if graph[j][i] == "#"{
				count++
		    }
	    }
		fmt.Println(count)
		count = 0
	}
}
