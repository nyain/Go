package main

import "fmt"

func main() {
  var uang int
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")

  fmt.Print("Masukkan uang: ")
  fmt.Scan(&uang)

for uang % 25 != 0 {
  fmt.Println(uang," Tidak valid")
  fmt.Print("Masukkan uang: ")
  fmt.Scan(&uang)
}
  fmt.Println(uang," Valid")
}
