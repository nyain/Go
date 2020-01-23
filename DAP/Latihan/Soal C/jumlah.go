package main

import "fmt"

func main() {
  var n, Sum int

  fmt.Scanln(&n)

  Sum = 0
  for i :=0; i<n; i++{
    Sum += (i+1)
    fmt.Println(i+1)
  }
  fmt.Println(Sum)
}
