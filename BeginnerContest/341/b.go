package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var N, tmp int
	var A []int
	fmt.Fscan(r, &N)
	for i := 0; i < N; i++ {
		fmt.Fscan(r, &tmp)
		A = append(A, tmp)
	}

	var S, T int
	for i := 0; i < N-1; i++ {
		fmt.Fscan(r, &S, &T)
		var times = A[i] / S
		A[i+1] += T * times
	}

	fmt.Print(A[len(A)-1])
}
