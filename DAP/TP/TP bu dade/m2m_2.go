package main
import "fmt"
func main() {
  var kkm,n,i,nilai,jumlah,rata1 float64
  var menang bool
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")
  fmt.Scan(&kkm)
  fmt.Scan(&n)
  for n != 0 {
    n *= 3
    i = 0
    for i < n{
      fmt.Scan(&nilai)
      jumlah = jumlah + nilai
      rata1 = jumlah / n
      menang = rata1 >= kkm
      i++
    }
    fmt.Println("Rata-rata peserta adalah ",rata1,"Menang ?",menang)
    jumlah = 0
    rata1 = 0
    n = 0
    fmt.Scan(&n)
  }
}
