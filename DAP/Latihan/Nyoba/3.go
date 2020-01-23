package main

import "fmt"

func main() {
  var a,b,c,d,e,f,r1,r2 float64
  var cek bool

  fmt.Scanln(&a,&b,&c) //Ahmad
  fmt.Scanln(&d,&e,&f) //Badrun

  r1=(a+b+c)/3 //rata-rata ahmad
  r2=(d+e+f)/3 //rata-rata badrun
  cek = r1>r2
  fmt.Println(r1,r2)
  fmt.Println("Apakah nilai Ahmad lebih baik dari nilai Badrun?",cek)
}
