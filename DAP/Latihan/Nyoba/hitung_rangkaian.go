package main

import "fmt"

func main() {
  var r1,r2,r3,rs float64
  var rp float64

  fmt.Print("Masukkan hambatan 1: ")
  fmt.Scanln(&r1)
  fmt.Print("Masukkan hambatan 2: ")
  fmt.Scanln(&r2)
  fmt.Print("Masukkan hambatan 3: ")
  fmt.Scanln(&r3)

  rs=r1+r2+r3
  rp= 1/r1 + 1/r2 + 1/r3
  rp= 1/rp
  fmt.Println("Rangkaian seri: ", rs)
  fmt.Println("Rangkaian pararel: ", rp)

}
