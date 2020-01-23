package main
import (
  "fmt"
)
func main() {
  var (
    s,i,n,x,m float64
    j int64
  )
  x=1
  m=1
  fmt.Print("N suku pertama : ")
  fmt.Scan(&n)

  for i < n{

      s=s+(1/x) //1-1/3+1/5-1/7
      m=m+2 //3 5 7 9

      if j%2 == 1 {
        m = -m //-5 -9  //minus
      }

      x = -m //-3 5 -7

      if j%2 == 1 {
        m = -m
      }

      j++ //0 1 2 3
      i++
  }
  fmt.Println(s*4)
  fmt.Println(i)
}
