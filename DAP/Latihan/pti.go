package main
import "fmt"
func main() {
  var n,a,b,c,harga,i float64
  fmt.Print("Harga IPO: ")
  fmt.Scan(&n)
  for i = 1; i < 6; i++ {
    fmt.Print("Harga minggu ke-",i,": ")
    fmt.Scan(&harga)
    if harga > n {
      fmt.Println("Status BULLISH")
      a++
      n = harga
    }else if harga < n {
      fmt.Println("Status BEARISH")
      b++
      n = harga
    }else if harga == n {
      fmt.Println("Status SIDEWAYS")
        c++
    }
  }
  fmt.Println("Terjadi ",a," BULLISH,",b," BEARISH, dan",c," SIDEWAYS")
  fmt.Println("Harga rata-rata: ",n/6)
}
