package main
import "fmt"
func main() {
  var n,i,j,nilai,peserta,jumlah,rata2,rata2keseluruhan,kkm,hasilakhir float64
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")

  fmt.Scan(&peserta,&n,&kkm)
  n = n*3
  i = 0
  for i < peserta {
      i++

      j=0
      jumlah = 0
      for j < n {
        fmt.Scan(&nilai)
        jumlah = jumlah + nilai
        j++
      }
      rata2 = jumlah/n
      menang := rata2 >= kkm
      fmt.Println("Peserta",i,"menang ", menang)
      fmt.Println("Rata-rata peserta",i, rata2)
      rata2keseluruhan = rata2keseluruhan + rata2
  }
  hasilakhir = rata2keseluruhan/peserta
  cek := hasilakhir >= kkm
  fmt.Println("Apakah rata-rata seluruh peserta melampau kkm? ", "(",hasilakhir,")", cek)
}
