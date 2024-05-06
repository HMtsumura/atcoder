package main

import "fmt"

func main() {
	var W, B int
	fmt.Scan(&W, &B)
	var countW = W / 7
	var restW = W - countW*7
	var restB = B - countW*5
	if restW-restB >= 3 {
		fmt.Print("No")
	} else {
		fmt.Print("Yes")
	}
}
