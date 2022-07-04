package main

import (
	"fmt"
	"time"
)

func main(){
	var n int

	fmt.Scan(&n)

	t :=  time.Date(0, 0, 21, 0, 0, 0, 0, time.UTC)
	t2 := t.Add(time.Duration(n)*time.Minute)

	fmt.Printf("%02d:%02d", t2.Hour(), t2.Minute())
}