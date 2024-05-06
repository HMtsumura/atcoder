package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var listA []int
	position := make([]int, N)
	for i := 0; i < N; i++ {
		var A int
		fmt.Scan(&A)
		listA = append(listA, A)
		position[A-1] = i + 1
	}
	fmt.Println(position)
	list := make(map[int]int, N)
	for i := 0; i < N; i++ {
		if listA[i] == i+1 {
			continue
		} else {
			fmt.Print(position[i])
			tmp1 := listA[position[i]-1]
			tmp2 := listA[i]
			listA[i] = tmp2
			listA[position[i]-1] = tmp1
			list[tmp1] = tmp2
		}
	}

	fmt.Println(len(list))
	for k, v := range list {
		fmt.Println(k, "", v)
	}
	// var count int
	// list := make(map[int]int)
	// for k, v := range listA {
	// 	if listA[k] != k+1 {
	// 		tmp := listA[k]

	// 		list[tmp] = listA[v-1]
	// 		listA[k] = listA[v-1]
	// 		listA[v-1] = tmp
	// 		count++
	// 	}
	// }

	// fmt.Println(count)
	// for k, v := range list {
	// 	fmt.Println(k, "", v)
	// }
}
