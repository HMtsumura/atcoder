package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	A := make(map[int]int)
	minSum := 0
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		if A[a] != 1 && a <= K {
			A[a] = 1
			minSum += a
		}
	}
	var sum int
	for i := 1; i <= K; i++ {
		sum += i
	}
	fmt.Print(sum - minSum)
}
