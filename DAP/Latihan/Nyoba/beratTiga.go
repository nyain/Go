package main
import "fmt"
func main() {
  var berat,n,i,kelompok1,kelompok2,kelompok3,kelompok4 float64
  fmt.Scan(&n)
  for i = 0; i < n; i++{
  fmt.Print("Masukkan berat: ")
  fmt.Scan(&berat)

  if berat >= 45 && berat <= 70 {
    kelompok1 += 1
  }else if berat > 70 && berat <= 100 {
    kelompok2 += 1
  }else if berat < 45 {
    kelompok3 += 1
  }else {
    kelompok4 +=1
  }
  }
  fmt.Println("Dari ",n,"orang dikelompokkan jadi: ")
  fmt.Println("Standar:",kelompok1)
  fmt.Println("Berat:",kelompok2)
  fmt.Println("Kekurangan gizi:",kelompok3)
  fmt.Println("Obesitas:",kelompok4)
}
