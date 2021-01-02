package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var h, w, minimum int
	var num_list []int
	//h,wの受け取り
	fmt.Scanf("%d %d", &h, &w)
	//
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < h; i++ {
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\n")
		numbers := strings.Split(text, " ")

		for j, val := range numbers {
			num, _ := strconv.Atoi(val)
			if (i == 0 && j == 0) || num < minimum {
				minimum = num
			}
			num_list = append(num_list, num)
		}
	}
	diff_sum := 0
	for _, number := range num_list {
		diff_sum += number - minimum
	}
	fmt.Println(diff_sum)
}
