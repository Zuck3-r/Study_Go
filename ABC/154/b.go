package main

import(
	"fmt"
	"strings"
)

func main() {
	var s, ans string

	fmt.Scan(&s)
	ans = strings.Repeat("x", len(s))
	fmt.Println(ans)
}
