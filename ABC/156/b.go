package main

import(
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var ans int
	for n > 0 {
		n /= k
		ans++
	}
	fmt.Println(ans)
}