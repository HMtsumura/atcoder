package main

import (
	"fmt"
)

func main() {
	var N, p, Q int
	fmt.Scan(&N)
	var P = make(map[int]int)
	for i := 0; i < N; i++ {
		fmt.Scan(&p)
		P[p] = i
	}
	var A, B int
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		fmt.Scanf("%d %d", &A, &B)
		if P[A] > P[B] {
			fmt.Println(B)
		} else {
			fmt.Println(A)
		}
	}
}
