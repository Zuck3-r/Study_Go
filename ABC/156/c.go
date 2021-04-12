package main

import(
	"fmt"
)

func main(){
	var n int
	fmt.Scan(&n)

	ary := make([]int, n)
	for i := range ary { fmt.Scan(&ary[i]) }

	ans := 100000000
	for p := 1; p < 100; p++ {
		sum := 0
		for _, x := range ary {
			sum += (x-p) * (x-p)
		}
		if sum < ans { ans = sum }
	}
	fmt.Println(ans)
}