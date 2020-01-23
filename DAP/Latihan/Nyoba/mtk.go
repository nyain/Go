package main

import "fmt"

func main() {
  var angka1,angka2 int
  var banding bool


  fmt.Print("Masukkan angka pertama: ")
  fmt.Scan(&angka1)

  fmt.Print("Masukkan angka kedua: ")
  fmt.Scan(&angka2)

  banding = angka1>angka2

  fmt.Println("Apakah angka pertama lebih besar dari angka kedua? ",banding)
}
