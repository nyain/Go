package main
import "fmt"

func main() {
  var A, B, Result int

  fmt.Scanln(&A, &B)

  Result = A-B

  if Result<0{
    Result *= -1
  }
  fmt.Println(Result)
}
