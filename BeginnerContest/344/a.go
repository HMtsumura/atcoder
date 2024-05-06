package main

import (
	"fmt"
	"regexp"
)

func main() {
	var S string

	fmt.Scanf("%s", &S)
	re := regexp.MustCompile("\\|.*\\|")
	result := re.ReplaceAllString(S, "")
	fmt.Print(result)

}
