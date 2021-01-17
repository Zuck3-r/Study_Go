package main

import (
	"fmt"
)

func main() {
	week := map[string]int{"SUN": 7, "MON": 6, "TUE": 5, "WED": 4, "THU": 3, "FRI": 2, "SAT": 1}
	var S string
	fmt.Scan(&S)
	fmt.Println(week[S])
}
