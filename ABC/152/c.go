package main

import(
	"fmt"
)

fucn main(){
	var n, t, min int
	c := 1
	fmt.Scan(&n)
	fmt.Scan(&min)

	for i := 1; i < n; i++ {
		fmt.Scan(&t)
		if t < min {
			c++
			min = t
		}
	}
	fmt.Println(c)
}
