package main

import (
	"fmt"
)

func main() {
	var A int
	var list []int
	for true {
		fmt.Scan(&A)
		list = append(list, A)
		if A == 0 {
			break
		}
	}
	for i := 0; i < len(list); i++ {
		fmt.Println(list[len(list)-1-i])
	}
}
