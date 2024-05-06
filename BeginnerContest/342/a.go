package main

import (
	"fmt"
)

func main() {
	var S string

	fmt.Scan(&S)
	var target byte
	if S[0] == S[1] {
		target = S[0]
	} else if S[0] != S[1] && S[1] != S[2] {
		target = S[0]
	} else {
		target = S[1]
	}
	for i := 0; i < len(S); i++ {
		if S[i] != target {
			fmt.Print((i + 1))
			break
		}
	}
}
