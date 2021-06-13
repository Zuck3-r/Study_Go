package main

import "fmt"

func main(){
    var n, a, b int64
    fmt.Scanf("%d", &n)
    fmt.Scanf("%d", &a)
    fmt.Scanf("%d", &b)
    ans := n / (a + b) * a
    n %= a + b
    if n <= a {
        ans += n
    } else {
        ans += a
    }
    fmt.Println(ans)
}
