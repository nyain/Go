package main
import "fmt"
func main() {
  var n,max,nmax,bil int
  max = -9999
  nmax = 0
  fmt.Scanln(&n)
  for i := 0; i < n; i++ {
    fmt.Scanln(&bil)
    if bil > max {
      max = bil
      nmax = 1
    }else if bil == max {
      nmax += 1
    }
  }
  fmt.Println("terbesar: ",max)
  fmt.Println("berapa kali: ",nmax)
}
