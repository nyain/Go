package main

import "fmt"

func main() {
  var bil int

  fmt.Scanln(&bil)

  x:=0
  i:=0

  if bil%2==1 {
    bil++
  }

  for i<6{
    fmt.Println(bil+x)
    i++
    x+=2
  }
}
