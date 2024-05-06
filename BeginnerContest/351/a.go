package main

import (
	"fmt"
)

func main() {
	var A, B int
	var countA, countB int
	for i := 0; i < 9; i++ {
		fmt.Scan(&A)
		countA += A
	}

	for i := 0; i < 8; i++ {
		fmt.Scan(&B)
		countB += B
	}

	fmt.Print(countA - countB + 1)
}
