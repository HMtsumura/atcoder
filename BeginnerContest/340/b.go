package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var Q int
	var index, num int
	var A []int
	fmt.Fscan(r, &Q)
	for i := 0; i < Q; i++ {
		fmt.Fscan(r, &index, &num)
		if index == 1 {
			A = append(A, num)
		} else {
			fmt.Println(A[len(A)-num])
		}
	}
}
