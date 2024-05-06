package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scanf("%d", &N)

	var A = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 1; i <= N-1; i++ {
		fmt.Printf("%d ", A[i]*A[i-1])
	}
}
