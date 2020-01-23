package main
import "fmt"
func main() {
  var bil,nFaktor int
  fmt.Scanln(&bil)
  nFaktor = 0

  for i := 1; i <= bil; i++ {
      /*nFaktor adalah jumlah 0 dari pembagian(mod) perulangan bilangan contohnya:
        5 mod 1 = 0 -> nFaktor
        5 mod 2 = 1
        5 mod 3 = 2
        5 mod 4 = 1
        5 mod 5 = 0 -> nFaktor
      */

      //jika hasil mod = 0, maka nFaktor ditambahkan 1
      if bil % i == 0 {
        nFaktor += 1
      }
  }
  //jika penjumlahan nFaktor sama dengan 2, maka bilangan itu prima
  if nFaktor == 2 {
    fmt.Println("Prima")
  }else {
    fmt.Println("Bukan Prima")
  }
}
