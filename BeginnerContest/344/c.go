package main

import (
	"fmt"
)

func main() {
	var N, M, L, Q int
	fmt.Scan(&N)

	var A, B, C []int
	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		A = append(A, a)
	}
	fmt.Scan(&M)
	for i := 0; i < M; i++ {
		var b int
		fmt.Scan(&b)
		B = append(B, b)
	}

	fmt.Scan(&L)
	for i := 0; i < L; i++ {
		var c int
		fmt.Scan(&c)
		C = append(C, c)
	}

	distinct := make(map[int]int)

	for _, v1 := range A {
		for _, v2 := range B {
			for _, v3 := range C {
				distinct[v1+v2+v3] = 1
			}
		}
	}
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		var x int
		fmt.Scan(&x)
		// var ans = check(A, B, C, x)
		if distinct[x] == 1 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
