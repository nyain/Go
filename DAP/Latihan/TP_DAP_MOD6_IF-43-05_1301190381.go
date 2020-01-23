
package main
import "fmt"

func Functionf(x float64) float64 {
  var f float64
  f = x*x
  return f
}
func Functiong(x float64) float64 {
  var g float64
  g = x-2
  return g
}
func Functionh(x float64) float64 {
  var h float64
  h = x+1
  return h
}
func Functionfgh(x float64) float64 {
  var fogoh float64
  fogoh = Functionf(Functiong(Functionh(x)))
  return fogoh
}

func main() {
  var x float64
  fmt.Print("Masukkan nilai X = ")
  fmt.Scan(&x)
  fmt.Println(Functionf(x))
  fmt.Println(Functiong(x))
  fmt.Println(Functionh(x))
  fmt.Println(Functionfgh(x))
}
