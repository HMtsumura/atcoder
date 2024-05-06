package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	result := make(map[string]int)
	for i := 0; i < len(S); i++ {
		for j := i; j < len(S); j++ {
			if result[S[i:j+1]] != 1 {
				result[S[i:j+1]] = 1
			}
		}
	}
	fmt.Print(len(result))
}
