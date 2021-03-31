package main

import(
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	x := map[int]bool{}

	for i := 0; i < N; i++ {
		var a int
		fmt.Scan(&a)
		if x[a] {
			fmt.Println("NO")
			return
		}
		x[a] = true
	}
	fmt.Println("YES")
}
