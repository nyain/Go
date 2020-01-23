package main
import "fmt"
func main() {
  var berat float64
  fmt.Print("Masukkan berat: ")
  fmt.Scan(&berat)
  if berat >= 45 && berat <= 70 {
    fmt.Println("Standar")
  }else if berat > 70 && berat <= 100 {
    fmt.Println("berat")
  }else {
    fmt.Println("Darurat")
  }
}
