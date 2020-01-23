package main
import "fmt"
func main() {
  var a1,a2,a3,b1,b2,b3,ra,rb float64

  fmt.Print("Masukkan nilai Ahmad: ")
  fmt.Scan(&a1,&a2,&a3)

  fmt.Print("Masukkan nilai Badrun: ")
  fmt.Scan(&b1,&b2,&b3)
  //Tugas terbimbing
  ra = (a1+a2+a3)/3
  rb = (b1+b2+b3)/3
  //Tugas mandiri
  absolut1 := a1 > b1 && a2 > b2 && a3 > b3
  absolut2 := b1 > a1 && b2 > a2 && b3 > a3
  //Bonus Tantangan
  kompetisi1 := a1 > b1 && a2 > b2 && a3 > b3 || a1 > b1 && a2 > b2 && a3 < b3 ||
                a1 < b1 && a2 > b2 && a3 > b3 || a1 > b1 && a2 < b2 && a3 > b3

  kompetisi2 := a1 < b1 && a2 < b2 && a3 < b3 || a1 < b1 && a2 < b2 && a3 > b3 ||
                a1 > b1 && a2 < b2 && a3 < b3 || a1 < b1 && a2 > b2 && a3 < b3

  pemenang   := kompetisi1 == kompetisi2

  fmt.Println(ra,rb)
  fmt.Println("Rata-rata Ahmad lebih baik dari Badrun? ", ra>rb)
  fmt.Println("Apakah Ahmad pemenang absolut? ",absolut1)
  fmt.Println("Apakah Badrun pemenang absolut? ",absolut2)

  fmt.Println("Apakah Ahmad pemenang?",kompetisi1)
  fmt.Println("Apakah Badrun pemenang?",kompetisi2)
  fmt.Println("Tidak ada pemenang?",pemenang)
}
