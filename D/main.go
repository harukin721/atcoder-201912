package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// 0番目使わない
	counts := make([]int, N+1)

	// 出現回数をカウント
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		counts[num]++
	}

	var twice int   // 2回出現する数値
	var missing int // 1回も出現しない数値

	// 1からNまでの数値をチェック
	for i := 1; i <= N; i++ {
		if counts[i] == 2 {
			twice = i
		} else if counts[i] == 0 {
			missing = i
		}
	}

	if twice == 0 && missing == 0 {
		fmt.Println("Correct")
	} else {
		fmt.Println(twice, missing)
	}

}
