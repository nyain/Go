package main

import "fmt"

func main() {
  var n,n1,i,ahmad,badrun,totala,totalb float64
  fmt.Println("Nama: Vincent Williams Jonathan")
  fmt.Println("NIM: 1301190381")

  fmt.Print("Masukkan soal yang diisi Ahmad: ")
  fmt.Scan(&n)
  n *=3
  i = 0
  fmt.Print("Masukkan nilai Ahmad: ")

  for i < n{
    fmt.Scan(&ahmad)
    totala = totala + ahmad
    i++
  }

  fmt.Print("Masukkan soal yang diisi Badrun: ")
  fmt.Scan(&n1)
  n1 *=3
  fmt.Print("Masukkan nilai Badrun: ")

  for i = 0; i < n1; i++{
    fmt.Scan(&badrun)
    totalb = totalb + badrun
  }

  rata1 := totala/n
  rata2 := totalb/n1
  fmt.Println("Rata-rata Ahmad adalah: ",rata1,"Ahmad menang:",rata1>rata2)
  fmt.Println("Rata-rata Badrun adalah: ",rata2,"Badrun menang:",rata1<rata2)
}
