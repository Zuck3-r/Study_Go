package main

import(
	"fmt"
)

func main() {
	var N,A int
	fmt.Scan(&N)

	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		if A % 2 == 0 && A % 3 != 0 && A % 5 != 0 {
			fmt.Println("DENIED")
			return
		}
	}
	fmt.Println("APPROVED")
}
