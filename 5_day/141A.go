package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)
	b := map[string]string{"Sunny": "Cloudy", "Cloudy": "Rainy", "Rainy": "Sunny"}
	fmt.Print(b[S])
}
