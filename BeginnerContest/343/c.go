package main

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	var i = 1
	var ans int
	for {
		var tmp = i * i * i
		var tmpString = strconv.Itoa(tmp)
		var firstHalf, latterHalf string
		if len(tmpString)%2 == 0 {
			firstHalf = tmpString[:len(tmpString)/2]
			latterHalf = tmpString[len(tmpString)/2:]
		} else {
			firstHalf = tmpString[:len(tmpString)/2]
			latterHalf = tmpString[(len(tmpString)/2)+1:]
		}
		var reversedTmp = reverse(latterHalf)
		if firstHalf == reversedTmp && tmp <= N {
			ans = tmp
		}
		if tmp > N {
			fmt.Print(ans)
			return
		}

		i++
	}
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
