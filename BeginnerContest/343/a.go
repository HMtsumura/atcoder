package main

import (
	"fmt"
)

func main() {
	var A, B int

	fmt.Scan(&A, &B)
	for i := 0; i < 9; i++ {
		if A+B != i {
			fmt.Print(i)
			break
		}
	}
}
