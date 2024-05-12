package main

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	var original = K
	count := 1
	var A []int
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		A = append(A, a)
	}
	i := 0
	for {
		if K >= A[i] {
			K = K - A[i]
			i++
		} else {
			count++
			K = original
		}
		if i == N {
			break
		}
	}
	fmt.Print(count)
}
