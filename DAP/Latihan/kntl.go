package main
import "fmt"
import "math"
func main() {
  var bil float64
  fmt.Scanln(&bil)
  q := 0.5*(1+2.236)
  for bil >= 0 {
    Si := (1/2.236*math.Pow(q,bil)-math.Pow(1-q,bil))
    fmt.Println(Si)
    fmt.Scanln(&bil)
  }
  fmt.Println("Bilangan negatif")
}
