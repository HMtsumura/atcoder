package main

import (
	"fmt"
	"strings"
)

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)
	sList := strings.Split(strings.ToUpper(S), "")
	tList := strings.Split(T, "")
	var count int
	var n = -1
	for tK, tV := range tList {
		if tK == 2 && tV == "X" {
			// count++
			break
		}
		for sK, sV := range sList {
			if tV == sV && n < sK {
				n = sK
				// sList = append(sList[:sK], sList[sK+1:]...) すでに使ったものを消す必要がなかった
				count++
				break
			}
		}
	}
	if (tList[2] == "X" && count == 2) || (tList[2] != "X" && count == 3) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
