package main

import(
  "fmt"
  "math"
)

func main() {
  var N float64
  fmt.Scan(&N)
  fmt.Println(math.Ceil(N/2))
}
