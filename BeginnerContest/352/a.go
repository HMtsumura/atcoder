package main

import (
	"fmt"
)

func main() {
	var N, X, Y, Z int
	fmt.Scanf("%d %d %d %d", &N, &X, &Y, &Z)
	if X < Y {
		for i := X; i < Y; i++ {
			if i == Z {
				fmt.Print("Yes")
				return
			}
		}
	} else if X > Y {
		for i := X; i > Y; i-- {
			if i == Z {
				fmt.Print("Yes")
				return
			}
		}
	}
	fmt.Print("No")

}
