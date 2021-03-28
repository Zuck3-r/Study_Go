package main

import(
	"fmt"
	"sort"
)

func main(){
	var N, K int
	fmt.Scan(&N, &K)

	H := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&H[i])
	}
	sort.Slice(H, func(i, j int) bool {return H[i] > H[j]})

	ans := 0
	if N <= K {
		fmt.Print(ans)
	} else if K == 0 {		
		for _, s := range H {
			ans += s
		}
		fmt.Print(ans)
	} else {
		for _, s := range H[N-K-1:] {
			ans += s
		}
		fmt.Print(ans)
	}
}
