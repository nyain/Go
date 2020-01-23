package main
import "fmt"
func main() {
  var barang1,barang2 float64
  fmt.Print("Masukkan berat di kedua kantong: ")
  fmt.Scanln(&barang1,&barang2)

  for (barang1 + barang2 < 150) && (barang1 > 0 && barang2 > 0){
    fmt.Println("Sepeda akan oleng: ",barang1-barang2 >= 9 || barang2-barang1 >= 9)
    fmt.Print("Masukkan berat di kedua kantong: ")
    fmt.Scanln(&barang1,&barang2)
  }
  fmt.Println("Program selesai")
}
