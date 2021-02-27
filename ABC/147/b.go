package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scanf("%s", &S)

	count := 0
	for i := 0; i < len(S)/2; i++ {
		if S[i] != S[len(S)-i-1] {
			count++
		}
	}
}
