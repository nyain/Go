/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat fungsi untuk melakukan perulangan sebanyak yang user input. perulangan akan
    berhenti jika  input user lebih dari sama dengan 9

*/
package main
import "fmt"
func main() {
  var kiri,kanan float64
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")
  lanjut := false

  for !lanjut {
      fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
      fmt.Scanln(&kiri,&kanan)
      //tidak lanjut
      lanjut = kiri >= 9.0 || kanan >= 9.0
  }
  fmt.Println("Proses selesai (Berat melebihi kapasitas)")
}
