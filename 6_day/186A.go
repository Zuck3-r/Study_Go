package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m float64
	fmt.Scan(&n, &m)
	//結果の出力
	fmt.Println(math.Floor(n / m))
}
