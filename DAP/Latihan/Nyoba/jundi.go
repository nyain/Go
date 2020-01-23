package main
import "fmt"
func main() {
  var n,u1,u2,u3 int

  fmt.Scanln(&n)
  u1 = 1
  u2 = 1
  fmt.Println(u1)
  fmt.Println(u2)

  for i := 1; i < n; i++ {
    u3 = u1+u2
    fmt.Println(u3)
    u1 = u2
    u2 = u3
  }
}
