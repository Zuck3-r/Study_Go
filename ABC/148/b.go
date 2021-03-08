package main

import (
	"fmt"
	"time"
)

func main(){
	var a int
	var b,c string
	fmt.Scan(&a, &b, &c)
	
	for i := 0; i < a; i++{
		time.Sleep(time.Second * 2)
		fmt.Printf("%c%c", b[i], c[i])
	}
	// fmt.Println()
}
