package main

import(
	"fmt"
)

func main() {
	var A, B, K int
	fmt.Scan(&A, &B, &K)
	
	if A >= K {
	fmt.Print(A-K, B)
	} else if A+B-K <= 0 {
	fmt.Print(0, 0)
	} else {
	fmt.Print(0, A+B-K)
	}
}
