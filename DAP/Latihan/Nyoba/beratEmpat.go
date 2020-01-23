package main
import "fmt"
func main() {
  var berat,kelompok1,kelompok2,kelompok3 float64

  for berat != 9999 {
    fmt.Print("Masukkan berat: ")
    fmt.Scan(&berat)

    if berat >= 45 && berat <= 70 {
      kelompok1 += 1
    }else if berat > 70 && berat <= 100 {
      kelompok2 += 1
    }else if berat > 100 && berat != 9999 {
      kelompok3 += 1
    }

  }
  fmt.Println("\nData kelompok: ")
  fmt.Println("Standar:",kelompok1)
  fmt.Println("Berat:",kelompok2)
  fmt.Println("Obesitas:",kelompok3)
}
