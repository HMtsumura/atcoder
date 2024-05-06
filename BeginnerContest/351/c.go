package main

import (
	"fmt"
)

func main() {
	var N int

	fmt.Scan(&N)
	var A []int
	for i := 0; i < N; i++ {
		var tmp int
		fmt.Scan(&tmp)
		A = append(A, tmp)
		for {
			if len(A) <= 1 {
				break
			} else if A[len(A)-1] != A[len(A)-2] {
				break
			}
			var new = A[len(A)-1] + 1
			A = A[:len(A)-2]
			A = append(A, new)
		}
	}

	fmt.Println(A)
}
