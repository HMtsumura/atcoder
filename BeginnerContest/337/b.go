package main

import (
	"fmt"
	"strings"
)

const (
	A = 1
	B = 2
	C = 3
)

func main() {
	var s string
	fmt.Scan(&s)
	var a, b, c int
	a = strings.Count(s, "A")
	b = strings.Count(s, "B")
	c = strings.Count(s, "C")
	if strings.Repeat("A", a)+strings.Repeat("B", b)+strings.Repeat("C", c) == s {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
