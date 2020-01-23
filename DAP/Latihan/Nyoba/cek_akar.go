package main

import "fmt"

func main() {
  var a,b,c int
  var cek bool
  fmt.Print("Masukkan angka pertama: ")
  fmt.Scan(&a)
  fmt.Print("Masukkan angka kedua: ")
  fmt.Scan(&b)
  fmt.Print("Masukkan angka ketiga: ")
  fmt.Scan(&c)
  cek = b*b -4*a*c >= 0
  fmt.Println("Apakah persamaan memiliki akar? ", cek)

}
