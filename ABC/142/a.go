package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(float64(a/2+a%2) / float64(a))
}
