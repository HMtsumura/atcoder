package main

import {
	"fmt"	
}


func main() {
	var S string
	fmt.Scan(&S)
	result := make([]string, 1, 2)
	for i := 0; i < len(S); i++ {
		for j := i; j < len(S[i:]); j++ {
			if !	slices.Contains(result, S[i:j]) {
				result.append(S[i:j])
			}
		}
	}
	fmt.Print((result))
}
