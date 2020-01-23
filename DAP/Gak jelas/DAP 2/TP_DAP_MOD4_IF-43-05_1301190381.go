/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini berisi fungsi untuk mencari suatu nilai pi.
    Pi tersebut dicari dengan rumus Leibniz(Leibniz Formula).
    Program menggunakan konsep perulangan dan percabangan
    Program akan berhenti jika selisih dari kedua pi lebih dari sama dengan 0.00001
*/
package main
import "fmt"

func main() {
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381\n")
  var (
    a,b,x,y float64 // a = pi1,y = pi2, x = penyebut(wajib ganjil), p = minus positif, b = suku ke-n
    p int
  )
  x=1
  p=1
  a=0
  y=1

  for 4*a - 4*y >= 0.00001 || 4*y - 4*a >= 0.00001 {
    y = a
    if p % 2 == 1 {
      a += 1/x
    } else {
      a -= 1/x
    }
    x += 2
    p++
    b++
  }
  fmt.Println("Hasil PI: ",a*4)
  fmt.Println("Hasil PI: ",y*4)
  fmt.Println("Pada n ke: ",b+1)
}
