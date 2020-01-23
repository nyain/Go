/*
   Nama : RIFKI PRATAMA KHOIR
   Nim  : 1303190001
   program ini dibuat untuk mencari pangkat dua dari suatu Bilangan
*/


package main

import "fmt"

func main() {
  var bil int

  fmt.Printf("Masukkan Bilangan: ")
  fmt.Scanln(&bil)

  bil = bil*bil

  fmt.Print("Hasilnya adalah: ")
  fmt.Println(bil)
  fmt.Scanln()
}
