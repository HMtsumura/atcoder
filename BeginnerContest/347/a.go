package main

import (
	"fmt"
)

func main() {
	var N int
	var K int
	fmt.Scanf("%d %d", &N, &K)

	var A = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < N; i++ {
		if A[i]%K == 0 {
			fmt.Printf("%d ", A[i]/K)
		}
	}
}
