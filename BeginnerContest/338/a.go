package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var S string
	fmt.Scan(&S)
	var first = S[0]
	var theRest = S[1:]
	if unicode.IsUpper(rune(first)) && theRest == strings.ToLower(theRest) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}

}
