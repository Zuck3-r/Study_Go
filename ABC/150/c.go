package main

import(
	"fmt"
	"math"
)

func factorial(a int) int {
	res := 1
	for a > 0 {
		res = res * a
		a--
	}
	return res
}

func dictionary(a []int) int{
	res := 1
	for i,_ := range a {
		c := 1
		for j := i+1; j < len(a) ; j++ {
			if a[i] > a[j] {
				c++
			}
		}
		res += factorial(len(a)-(i+1)) * (c-1)
	}
	return res
}

func main(){
	var n int
	p := make([]int, n)
	q := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(p[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(q[i])
	}

	pnum := dictionary(p)
	qnum := dictionary(q)
	fmt.Println(math.Abs(float64(pnum - qnum)))
}
