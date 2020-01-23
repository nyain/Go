package main
import "fmt"
func main() {
  var a,b float64
  lanjut := false

  for !lanjut {
  fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
  fmt.Scanln(&a,&b)
  total := a+b
  lanjut := total >= 9.0 || a == 9.0 == || b == 9.0 
  }
    fmt.Println("Proses selesai (Berat melebihi kapasitas)")
}
