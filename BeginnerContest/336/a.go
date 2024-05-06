package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print("L")
	for i := 0; i < N; i++ {
		fmt.Print("o")
	}
	fmt.Print("ng")
}
