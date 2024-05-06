package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print("1")
	for i := 0; i < N; i++ {
		fmt.Print("01")
	}
}
