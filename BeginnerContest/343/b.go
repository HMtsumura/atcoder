package main

import "fmt"

func main() {
	var N, A int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Scan(&A)
			if A == 1 {
				fmt.Print(j+1, " ")
			}
		}
		fmt.Println()
	}
}
