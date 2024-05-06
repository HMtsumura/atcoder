package main

import (
	"fmt"
	"strings"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)
	var sList, tList []string
	sList = strings.Split(S, "")
	tList = strings.Split(T, "")
	var correctCount int
	for i := 0; i < len(T); i++ {
		if sList[correctCount] == tList[i] {
			fmt.Print(i+1, " ")
		}
	}
}
