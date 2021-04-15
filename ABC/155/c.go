package main

import(
	"fmt"
	"bufio"
	"os"
	"sort"
)

func nextStr() string {
	sc.Scan()
	return sc.Text()
}

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N int
	fmt.Scan(&N)

	var tmp string

	S := make(map(string), N)
	for i := 0; i < N; i++ {
		temp = nextStr()
		S[temp]++
	}

	var max int
	key := make([]string, N)
	var count int
	for i, v := range S {
		if max < v {
			max = v
		}
		key[count] = i
		count++
	}
	sort.Sort(sort.StringSlice(key))
 
	for i := 0; i < len(key); i++ {
		if max == S[key[i]] {
			fmt.Println(key[i])
		}
	}
}
