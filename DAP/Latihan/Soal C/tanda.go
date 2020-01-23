package main

import "fmt"

func main() {
  var (
    bil int
  )

  fmt.Scanln(&bil)
  if bil>0 {
    fmt.Println("Positif")
  } else if bil<0 {
    fmt.Println("Negatif")
  } else {
    fmt.Println("Nol")
  }
}
