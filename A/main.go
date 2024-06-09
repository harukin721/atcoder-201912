package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	// strconv.Atoi() は、文字列を整数に変換する関数
	if num, err := strconv.Atoi(S); err == nil {
		fmt.Println(num * 2)
	} else {
		fmt.Println("error")
	}
}
