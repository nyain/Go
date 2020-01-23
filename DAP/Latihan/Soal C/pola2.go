package main

import "fmt"

func main() {
  pola2()
}

func pola2()  {
  var n int

  fmt.Scanln(&n)

  for j := 1; j<= n; j++{
    for i := n; i >= j; i--{
      fmt.Print(" ")
    }
    for k := 1; k <= j; k++{
      fmt.Print("* ")
    }
    fmt.Println()
  }

}
