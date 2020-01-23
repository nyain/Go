package main
import "fmt"
func main() {
  var a,b,c float64

  fmt.Print("Masukkan sisi : ")
  fmt.Scan(&a,&b,&c)
  if a == b && a == c {
    fmt.Println("Segitiga sama sisi")
  }else if a == b || a == c {
    if a != b || a != c {
      fmt.Println("Segitiga sama kaki")
    }
  }else {
    fmt.Println("Segitiga sembarang")
  }
}
