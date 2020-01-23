package main

import "fmt"

func main() {
  var detik,jam,menit int
  fmt.Print("Masukkan detik: ")
  fmt.Scan(&detik)
  jam= detik/3600
  detik= detik%3600

  menit= detik/60
  detik= detik%60

  fmt.Println("Jam: ",jam, "Menit: ",menit,"Detik: ",detik)
}
