package main

import "fmt"

func main() {
  var uang int
  var hasil10k,hasil2k,hasil5k int


  fmt.Print("Masukkan uang: ")
  fmt.Scan(&uang)
// {
  hasil10k = uang/10000
  uang = uang%10000 //buat masukin hasil ke yang terbesar, kalo dihapus dia jadi 2 hasil (silahkan dicoba kalo ngeyel)

  hasil5k = uang/5000
  uang = uang%5000

  hasil2k = uang/2000
  uang = uang%2000
                  // }
  fmt.Println("10 ribu ada: ",hasil10k," lembar")
  fmt.Println("5 ribu ada: ",hasil5k," lembar")
  fmt.Println("2 ribu ada: ",hasil2k," lembar")

}
