package main

import "fmt"

func main() {
	var S string
	R_cnt := 0
	Odd := false
	Even := false
	K_flag := false

	fmt.Scan(&S)

	for i := 0; i < len(S); i++ {
		if S[i] == 'R' {
			if R_cnt == 1 && !K_flag {
				fmt.Print("No")
				return
			}
			R_cnt++
		} else if S[i] == 'B' {
			if i%2 == 0 && !Even {
				Even = true
			} else if i%2 == 1 && !Odd {
				Odd = true
			}
		} else if S[i] == 'K' {
			if R_cnt == 0 {
				fmt.Print("No")
				return
			}
			K_flag = true
		}
	}
	if !Even || !Odd {
		fmt.Print("No")
		return
	}
	fmt.Print("Yes")
}
