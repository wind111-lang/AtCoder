package main

import (
	"fmt"
)

type Haiku struct {
	haiku string
}

type Result struct {
	haiku string
	score int
}

//m := make(map[string]bool)

// func unset(str []Haiku, index int) []Haiku {
// 	//fmt.Println("unset", str, index)
// 	return append(str[:index], str[index+1:]...)
// }

func main() {
	var n, t int
	var s string
	fmt.Scanf("%d", &n)

	set := make(map[Haiku]bool)
	//res := make(map[Result]bool)

	str_res := []Result{}
	flag := []bool{}

	//m := []Result{}

	max := 0
	index := 0

	for i := 0; i < n; i++ {
		fmt.Scanf("%s %d", &s, &t)

		if _, ok := set[Haiku{s}]; ok {
			//res[Result{s, t}] = true
			str_res = append(str_res, Result{s, t})
			flag = append(flag, true)
		} else {
			set[Haiku{s}] = true
			//res[Result{s, t}] = false
			str_res = append(str_res, Result{s, t})
			flag = append(flag, false)
		}
	}
	for i := 0; i < n; i++ {
		if str_res[i].score > max && flag[i] == false {
			max = str_res[i].score
			index = i + 1
		}
	}

	// fmt.Println(str_res)
	// fmt.Println(flag)

	fmt.Println(index)

	//max := 0

	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%s %d", &haiku[i].haiku, &haiku[i].score)
	// 	for j := 0; j < n; j++ {
	// 		//fmt.Println(i, j)
	// 		if (i != j) && (haiku[i].haiku == haiku[j].haiku) {
	// 			//fmt.Println("same", haiku[i].haiku, haiku[j].haiku, i, j)
	// 			haiku = append(haiku, Haiku{haiku: haiku[i].haiku, score: haiku[i].score, flag: true})
	// 			flag = true
	// 			break
	// 			//fmt.Println("true!")
	// 		}
	// 	}
	// 	if flag == false {
	// 		haiku = append(haiku, Haiku{haiku: haiku[i].haiku, score: haiku[i].score, flag: false})
	// 	}

	// 	if i != n {
	// 		haiku = unset(haiku, 0)
	// 	}
	// 	flag = false
	// 	fmt.Println(haiku)
	// }

	// fmt.Println(haiku)
	//var s string
	//fmt.Scan(&n)

	//max := 0
	//max_index := 1
	//index := make(map[int]bool)

	//cnt := 1

	//garbage := "z"

	//m := make(map[string]int)
	//check := make(map[string]bool)

	// for i := 0; i < n; i++ {
	// 	fmt.Scanf("%s %d", &s, &num)
	// 	if _, ok := m[s]; ok {
	// 		//m[] = num
	// 		m[garbage] = num
	// 		index[i] = true
	// 		garbage += garbage
	// 		continue
	// 	} else {
	// 		m[s] = num
	// 		index[i] = false
	// 	}
	// }

	// fmt.Println()

	// for _, v := range m {
	// 	//fmt.Println(k, v)
	// 	//if _, t := m[k]; t && index[k] == true {
	// 	//delete(m, k)
	// 	//fmt.Println("deleted", k, v)
	// 	//delete(m,v)
	// 	//}
	// 	if index[cnt-1] == true {
	// 		continue
	// 	} else {
	// 		if v > max && index[v] != true {
	// 			//fmt.Println("MAX!", v)
	// 			max_index = cnt
	// 			max = v
	// 		}
	// 	}
	// 	cnt++
	// }
	// fmt.Println(index)
	// fmt.Println(m)

	// fmt.Println(max_index)
}
