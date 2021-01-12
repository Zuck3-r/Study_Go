package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	a -= b * 2
	if a < 0 {
		a = 0
	}
	fmt.Println(a)
}
