package main

import (
	"fmt"
	"sort"
)

func main() {
	var s string
	fmt.Scan(&s)
	var mapS = make(map[string]int)
	for _, v := range s {
		mapS[string(v)]++
	}
	var keys []string
	for k, _ := range mapS {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var max int
	var maxString string
	for _, v := range keys {
		if mapS[v] > max {
			max = mapS[v]
			maxString = v
		}
	}
	fmt.Print(maxString)
}
