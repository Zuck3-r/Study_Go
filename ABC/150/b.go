package main

import(
	"fmt"
	"strings"
)

func main(){
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)

	fmt.Printf("%d\n", strings.Count(s, "ABC"))
}
