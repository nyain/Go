/*
  Nama : RIFKI PRATAMA KHOIR
  Nim  : 1303190001
  program ini dibuat untuk mencari nilai voltase dari suatu
  rangkaian listrik
*/

package main

import "fmt"

func main() {
  var R,V,I int

  fmt.Println("Masukkan arus dan tahanan")
  fmt.Print("Arus: ")
  fmt.Scanln(&I)
  fmt.Print("Tahanan: ")
  fmt.Scanln(&R)

  V = I*R


  fmt.Println("Voltase: ",V)
  fmt.Scanln()

}
