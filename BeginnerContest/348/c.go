package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)
	beans := make(map[int]int)
	for i := 0; i < N; i++ {
		var A, C int
		fmt.Scan(&A, &C)

		if beans[C] == 0 || beans[C] > A {
			beans[C] = A
		}
	}
	var max int
	for _, v := range beans {
		if max == 0 || v > max {
			max = v
		}
	}
	fmt.Print(max)
}
