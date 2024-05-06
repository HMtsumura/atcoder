package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	var list = make(map[byte]int)
	for i := 0; i < len(S); i++ {
		list[S[i]]++
	}

	countSet := make(map[int]int)
	for _, count := range list {
		countSet[count]++
	}

	for _, count := range countSet {
		if count != 0 && count != 2 {
			fmt.Print("No")
			return
		}

	}
	fmt.Println("Yes")
}
