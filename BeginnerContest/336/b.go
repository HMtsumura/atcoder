package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var binary = fmt.Sprintf("%b", N)
	var count int
	if string(binary[len(binary)-1]) == "1" {
		fmt.Print(0)
		return
	}

	for i := len(binary) - 1; i >= 0; i-- {
		if string(binary[i]) == "0" {
			count++
		} else {
			fmt.Print(count)
			return
		}
	}
}
