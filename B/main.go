package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// N個の要素を持つスライスを作成
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	// 2日目以降
	for i := 1; i < N; i++ {
		if A[i] > A[i-1] {
			fmt.Printf("up %d\n", A[i]-A[i-1])
		} else if A[i] < A[i-1] {
			fmt.Printf("down %d\n", A[i-1]-A[i])
		} else {
			fmt.Printf("stay\n")
		}
	}
}
