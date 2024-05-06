package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	sumX := 0
	sumY := 0
	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		sumX += X
		sumY += Y
	}

	if sumX < sumY {
		fmt.Print("Aoki")
	} else if sumX > sumY {
		fmt.Print("Takahashi")
	} else {
		fmt.Print("Draw")
	}
}
