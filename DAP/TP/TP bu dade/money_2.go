package main

import "fmt"

func main() {
  var uang,i int
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")

  fmt.Print("Masukkan uang: ")
  fmt.Scan(&uang)

for uang >= 10000 {
  uang -= 10000
  i++
}
  fmt.Println("Pecahan 10000 ada : ",i," lembar")
}
