/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat konsep pemograman modular dengan membuat fungsi dan prosedur
    Terdapat 4 fungsi di sini di mana masing-masing fungsi tersebut memiliki algoritma yang berbeda-beda
    fungsi 1: membuat operasi aritmatika kuadrat dari suatu nilai variabel
    fungsi 2: membuat operasi aritmatika yaitu mengurang suatu nilai variabel dengan 2
    fungsi 3: membuat operasi aritmatika yaitu menambah suatu nilai variabel dengan 1
    fungsi 4: fungsi irisan dari ketiga 3 fungsi di atas

*/
package main
import "fmt"

func f(x float64) float64 {
  var f float64
  f = x*x
  return f
}
func g(x float64) float64 {
  var g float64
  g = x-2
  return g
}
func h(x float64) float64 {
  var h float64
  h = x+1
  return h
}
func fogoh(x float64) float64 {
  var fogoh float64
  fogoh = f(g(h(x)))
  return fogoh
}

func main() {
  var x float64
  fmt.Print("Masukkan nilai X = ")
  fmt.Scan(&x)
  fmt.Println("f(",x,") = ",f(x))
  fmt.Println("g(",x,") = ",g(x))
  fmt.Println("h(",x,") = ",h(x))
  fmt.Println("fogoh(",x,") =",fogoh(x))
}
