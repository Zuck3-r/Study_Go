package main

import(
	"fmt"
)

func main(){
	var a, b int
	fmt.Scan(&a, &b)

	if b < a {
		a, b =b, a
	}
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
	fmt.Println()
}
