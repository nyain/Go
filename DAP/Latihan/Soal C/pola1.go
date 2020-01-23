package main

import "fmt"

func main() {
  var i,j,n int

  fmt.Scanln(&n)

  for i=1 ; i<=n; i++ {
        for j=i; j<n; j++ {
            fmt.Print(" ") // segitiga kiri
        }
        for j=1; j<=i; j++ {
            fmt.Print("*") // segitiga kanan
        }
        fmt.Println()
    }
}
