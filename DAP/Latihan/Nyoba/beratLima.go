package main
import "fmt"
func main() {
  var berat,kelompok1,kelompok2,kelompok3,totalstandar,totalberat,totalobesitas,rata1,rata2,rata3 float64

  for berat != 9999 {
    fmt.Print("Masukkan berat: ")
    fmt.Scan(&berat)

    if berat >= 45 && berat <= 70 {
      kelompok1 += 1
      totalstandar = totalstandar + berat
    }else if berat > 70 && berat <= 100 {
      kelompok2 += 1
      totalberat = totalberat + berat
    }else if berat > 100 && berat != 9999 {
      kelompok3 += 1
      totalobesitas = totalobesitas + berat
    }
    rata1 = totalstandar / kelompok1
    rata2 = totalberat / kelompok2
    rata3 = totalobesitas / kelompok3
  }
  fmt.Println("Data kelompok: ")
  fmt.Println("Standar:",kelompok1)
  fmt.Println("Berat:",kelompok2)
  fmt.Println("Obesitas:",kelompok3)
  fmt.Println(rata1) //standar
  fmt.Println(rata2) //berat
  fmt.Println(rata3) //obesitas
}
