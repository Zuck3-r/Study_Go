package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	ar := make([]int, N)
	for i := range ar {
		fmt.Scan(&ar[i])
	}
	ans := 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			ans += ar[i] * ar[j]
		}
	}
	fmt.Println(ans)
}
