/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat fungsi untuk melakukan perulangan sebanyak yang user input. perulangan akan
    berhenti jika total kedua input user lebih dari 150 dan juga jika kedua input user negatif.
    perulangan akan terus berjalan jika selisih kedua input user lebih dari sama dengan 9, namun
    akan menampilkan boolean false atau true.

*/
package main
import "fmt"
func main() {
    var kiri,kanan float64
    fmt.Println("Nama: Vincent Williams Jonathan")
    fmt.Println("NIM: 1301190381")
    fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
    fmt.Scanln(&kiri,&kanan)
    /*perulangan berhenti jika:
     1. berat lebih dari 150
     2. beban kiri dan kanan negatif
     */
    for (kiri+kanan < 150) && (kiri > 0 && kanan > 0) {
      //selisih
      selisih := kiri-kanan
      selisih2 := kanan-kiri
      /*menampilkan boolean true jika selisih lebih dari 9,
      tapi tidak memberhentikan looping */
      lanjut2 := selisih >= 9.0 || selisih2 >= 9.0
      fmt.Println("Sepeda motor pak Andi oleng: ",lanjut2)
      //input ulang
      fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
      fmt.Scanln(&kiri,&kanan)
    }
    fmt.Println("Proses selesai (Berat melebihi kapasitas)")
}
