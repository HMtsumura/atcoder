package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var H int
	var W int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	fmt.Scanf("%d %d", &H, &W)

	square := make([][]int, H)
	for i := 0; i < H; i++ {
		square[i] = make([]int, W)
		for j := 0; j < W; j++ {
			fmt.Fscan(reader, &square[i][j])
		}
	}

	sumRows := make([]int, H)
	sumCols := make([]int, W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			sumRows[i] += square[i][j]
			sumCols[j] += square[i][j]
		}
	}

	for i := 0; i < H; i++ {

		fmt.Fprintf(writer, "%d ", sumCols[i])
	}
	fmt.Fprintln(writer)
	// i行j列目に対してその十字の総和を求める
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			fmt.Fprintf(writer, "%d ", sumRows[i]+sumCols[j]-square[i][j])
		}
		fmt.Fprintln(writer)
	}
}
