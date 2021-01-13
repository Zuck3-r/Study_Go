package main

import (
	"fmt"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	if A > 9 || B > 9 {
		fmt.Println(-1)
		return
	}
	fmt.Println(A * B)
}
