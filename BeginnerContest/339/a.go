package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	var lastIndex = strings.LastIndex(S, ".")

	fmt.Print(S[lastIndex+1:])
}
