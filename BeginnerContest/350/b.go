package main

import (
	"fmt"
)

func main() {
	var N, Q int
	fmt.Scan(&N, &Q)
	teeth := make([]bool, N)
	for i := 0; i < Q; i++ {
		var position int
		fmt.Scan(&position)
		teeth[position-1] = !teeth[position-1]
	}
	var count int
	for _, v := range teeth {
		if !v {
			count++
		}
	}
	fmt.Print(count)
}
