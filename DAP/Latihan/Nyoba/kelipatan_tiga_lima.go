package main

import "fmt"

func main() {
  var  a int
  var cek bool

  fmt.Print("Masukkan angka: ")
  fmt.Scanln(&a)
  cek = a%3 == 0 && a%5 == 0
  fmt.Println("Apakah",a, "termasuk kelipatan 3 dan 5? ", cek)
}
