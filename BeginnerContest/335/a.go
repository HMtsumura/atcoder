package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Print(S[:len(S)-1], "4")
}
