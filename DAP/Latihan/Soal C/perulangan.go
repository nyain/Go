package main
import "fmt"
func main() {
  var (
      s string
      n int
  )
  fmt.Scanln(&s,&n)

  for i := 0; i<n; i++ {
    fmt.Println(s)
	}
}
