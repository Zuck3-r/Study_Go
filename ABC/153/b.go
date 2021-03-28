package main
import "fmt"
func main(){
    var h, n, a int
    fmt.Scan(&h, &n)
    
    for i := 0; i < n; i++ {
        fmt.Scan(&a)
        h -= a
    }
    if h <= 0 {
        fmt.Print("Yes")
        return
    }
    fmt.Print("No")
    return
}
