package main

import (
	"fmt"
	"strconv"
)

func main(){
	var A, B, X int
	fmt.Scan(&A, &B, &X)

	var l, r int
	l = 0
	r = 1000000001

	var ans int
	var temp int = (l + r) / 2

	for {
		ans = A * temp + B * len(strconv.Itoa(temp))
		if ans <= X {
			l = temp
		} else {
			r = temp
		}
		temp = (l + r) / 2
		if temp == l {
			fmt.Println(temp)
			break
		}
	}
}
