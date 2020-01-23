package main

import "fmt"

func main() {
  var sukuawal, beda, n int

  fmt.Scanln(&n,&sukuawal,&beda)

  for i := 0; i<n; i++{
    fmt.Println(sukuawal + (beda * i))
  }
}
