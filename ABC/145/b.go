package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N)
	fmt.Scan(&S)
	if S[:N/2] == S[N/2:] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
