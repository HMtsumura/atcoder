package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var listA, listB []int

	var height int
	for i := 0; i < N; i++ {
		var A, B int
		fmt.Scanf("%d %d", &A, &B)
		listA = append(listA, A)
		listB = append(listB, B)
		height += A
	}
	var maxGap, gap, h int
	for i := 0; i < len(listA); i++ {
		gap = listB[i] - listA[i]
		if gap > maxGap {
			if i != 0 {
				h = h + listB[i] - listB[i-1] + listA[i-1]
			} else {
				h = h + listB[i]
			}
			maxGap = gap
		} else {
			h += listA[i]
		}
	}
	// var maxHeight, h int
	// for i := 0; i < len(listA); i++ {
	// 	h = height - listA[i] + listB[i]
	// 	if h > maxHeight {
	// 		maxHeight = h
	// 	}
	// }
	fmt.Print(h)
}
