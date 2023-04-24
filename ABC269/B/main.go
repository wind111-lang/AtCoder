package main

import(
	"fmt"
	"strings"
)

func main() {
	var line string

	var a,b,c,d int
	flag := false

	arr := make([][]string,0,10)

	for i:=0; i<10; i++{
		fmt.Scan(&line)
		arr = append(arr,strings.Split(line,""))
	}
	//fmt.Println(arr)

	for i:=0; i<10; i++{
		for j:=0; j<10; j++{
			if arr[j][i] == "#" && !flag{
				a = j+1
				flag = true
			}
			if j > 0{
			    if arr[j][i] == "." && arr[j-1][i] == "#" && flag{
				    b = j
				    break
			    }else{
				    b = 10
			    }
		    }
		}
		if flag{
			flag = false
			break
		}
	}
	for i:=0; i<10; i++ {
		for k:=0; k<10; k++{
			if arr[i][k] == "#" && !flag{
				c = k+1
				flag = true
			}
			if k > 0{
			    if arr[i][k] == "." && arr[i][k-1] == "#" && flag{
				    d = k
				    break
			    }else{
				    d = 10
			    }
		    }
		}
		if flag{
			flag = false
			break
		}
	}
	fmt.Println(a,b)
	fmt.Println(c,d)
}