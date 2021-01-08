package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	al := make([]int, n)
	for i := range al {
		fmt.Scan(&al[i])
	}
	bl := make([]int, n)
	for i := range bl {
		fmt.Scan(&bl[i])
	}
	cl := make([]int, n-1)
	for i := range cl {
		fmt.Scan(&cl[i])
	}
	ans := bl[al[0]-1]
	for i := 1; i < n; i++ {
		ans += bl[al[i]-1]
		if al[i-1]+1 == al[i] {
			ans += cl[al[i]-2]
		}
	}
	fmt.Println(ans)
}
