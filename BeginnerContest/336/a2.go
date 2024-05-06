package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print("L", strings.Repeat("o", N), "ng")
}
