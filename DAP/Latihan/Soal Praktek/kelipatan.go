package main
import (
  "fmt"
  "math"
)
func main() {
  var (
    angka,batas,pembagi,i float64
    euy bool
  )
  euy = false
  angka = 1
  i=1
  for !euy{
    pembagi = 0
    batas = (int)math.Sqrt(angka)
    if angka % batas == 0 {
      pembagi++
    }
    for i < batas {
      if angka % i == 0 {
        i++
      }
    }
    if pembagi >= 500 {
      euy = true
      fmt.Print("" + angka)
    }
    angka++
  }
  fmt.Print("euy")
}
