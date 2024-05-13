package main

import (
	"fmt"
)

const (
	North = iota
	East
	South
	West
)

func main() {
	var H, W, N int
	fmt.Scan(&H, &W, &N)
	result := make([][]string, H)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			result[i] = append(result[i], ".")
		}
	}
	var x, y int
	var direction = North
	for i := 0; i < N; i++ {
		if result[x][y] == "." {
			result[x][y] = "#"
			if direction == North {
				y += 1
				if y == W {
					y = 0
				}
				direction = East
			} else if direction == East {
				x += 1
				if x == H {
					x = 0
				}
				direction = South
			} else if direction == South {
				y -= 1
				if y < 0 {
					y = W - 1
				}
				direction = West
			} else {
				x -= 1
				if x < 0 {
					x = H - 1
				}
				direction = North
			}

		} else {
			result[x][y] = "."
			if direction == North {
				y -= 1
				if y < 0 {
					y = W - 1
				}
				direction = West
			} else if direction == East {
				x -= 1
				if x < 0 {
					x = H - 1
				}
				direction = North
			} else if direction == South {
				y += 1
				if y == W {
					y = 0
				}
				direction = East
			} else {
				x += 1
				if x == H {
					x = 0
				}
				direction = South
			}
		}
	}
	for _, v := range result {
		for _, v1 := range v {
			fmt.Print(v1)
		}
		fmt.Println("")
	}
}
