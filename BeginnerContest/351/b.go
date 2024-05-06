package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	var A, B string
	var listA = make([][]string, N)
	var listB = make([][]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		l := strings.Split(A, "")
		for j := 0; j < len(l); j++ {
			listA[i] = append(listA[i], l[j])
		}

	}
	for i := 0; i < N; i++ {
		fmt.Scan(&B)
		l := strings.Split(B, "")
		for j := 0; j < len(l); j++ {
			listB[i] = append(listB[i], l[j])
		}

	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if listA[i][j] != listB[i][j] {
				fmt.Print(i+1, j+1)
				return
			}
		}
	}
	// fmt.Println(listA)
}
