package main
import "fmt"
func main() {
  var berat float64
  fmt.Print("Masukkan berat: ")
  fmt.Scan(&berat)

  for berat <= 45 {
    fmt.Println("Kekurangan gizi")
    fmt.Print("Masukkan berat: ")
    fmt.Scan(&berat)
  }

  if berat >= 45 && berat <= 70 {
    fmt.Println("Standar")
  }else if berat > 70 && berat <= 100 {
    fmt.Println("Berat")
  }else {
    fmt.Println("Darurat")
  }
}
