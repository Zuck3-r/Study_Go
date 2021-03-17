package main

import(
	"fmt"
)

func main(){
	var x int
	fmt.Scan(&x)
	for !isPrime(x) {
		x++
	}
	fmt.Println(x)
}

func isPrime(a int)bool {
	if a < 2 {
		return false
	}

	for j := 2; j * j <= a; j++ {
		if a % j == 0 {
			return false
		}
	}
	return true
}


