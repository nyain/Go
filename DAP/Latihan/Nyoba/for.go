package main
import "fmt"
func main() {
  var a,b,c,d string
  var cek bool

  for i := 0;i < 5;i++ {
    fmt.Printf("Masukkan warna: ")
    fmt.Scanln(&a,&b,&c,&d)
    cek = (a == "merah") && (b == "kuning") && (c == "hijau") && (d == "ungu") && cek
  }

  fmt.Println("BERHASIL?", cek)
}
