package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var first int
	flag := true
	for i := 0; i < N; i++ {
		var H int
		fmt.Scan(&H)
		if i == 0 {
			first = H
		} else {
			if H > first {
				fmt.Print(i + 1)
				flag = false
				break
			}
		}
	}
	if flag {
		fmt.Print(-1)
		return
	}
}
