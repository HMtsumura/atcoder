package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	var numContest = S[3:]
	if numContest == "316" || numContest > "349" || numContest < "001" {
		fmt.Print("No")
	} else {
		fmt.Print("Yes")
	}
}
