package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	distinct := make(map[byte]int)
	// Sの中に何回同じアルファベットが出てきたかを記録する
	for i := 0; i < len(S); i++ {
		distinct[S[i]] += 1
	}
	// 重複を考えずにSの文字数で組み合わせを考えた場合に何通りになるかを取得する
	var sum = len(S) * (len(S) - 1) / 2
	var b = false

	for _, v := range distinct {
		if v != 1 {
			b = true
		}
		// 同じアルファゲットの数から、重複する文字列の数を引く※重複分が全て引かれてしまうので、1通りは後に追加する
		sum = sum - (v * (v - 1) / 2)
	}
	if b {
		fmt.Print(sum + 1)
	} else {
		fmt.Print(sum)
	}
}
