package main
import "fmt"

func main() {
  var celcius,reamur,fahrenheit,kelvin float64

  fmt.Scanln(&celcius)
  reamur = celcius * 4/5
  fahrenheit = celcius * 9/5 + 32
  kelvin = celcius + 273.15

  fmt.Println("Dalam Fahrenheit:", fahrenheit)
  fmt.Println("Dalam Reamur:",reamur)
  fmt.Println("Dalam Kelvin:", kelvin)
}
