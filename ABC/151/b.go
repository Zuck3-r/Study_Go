package main

import(
	"fmt"
)

func main(){
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	t := m * n

	for i := 0; i < n - 1; i++ {
		var a int
		fmt.Scan(&a)
		t -= a
	}
	if t < 0{
		fmt.Print(0)
	} else if t <= k {
		fmt.Print(t)
	} else {
		fmt.Print(-1)
	}
}

