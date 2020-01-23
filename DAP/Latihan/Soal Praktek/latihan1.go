package main
import "fmt"
func main() {
  var(
    sum,i int
    tiga,lima bool
  )
  tiga=false
  lima=false
  for i = 3; i < 1000; i++ {
    tiga = i % 3 == 0
    lima = i % 5 == 0
    if tiga || lima {
      sum += i
    }
  }
  fmt.Println(sum)
}
