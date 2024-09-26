package main

import (
	"fmt"
	"sort"
)

func main() {
	var A, B, C, D, E, F int
	fmt.Scan(&A, &B, &C, &D, &E, &F)

	numbers := []int{A, B, C, D, E, F}

	// 昇順
	// intSlice := sort.IntSlice(numbers)

	// 逆順
	// reverse := sort.Reverse(intSlice)

	// ソート
	// sort.Sort(reverse)

	// まとめるとこう
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))

	fmt.Println(numbers[2])
}
