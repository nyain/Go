package main

import "fmt"

func main() {
  var teks string

  fmt.Print("Masukkan kalimat: ")
  fmt.Scanln(&teks)

  for i :=0; i < len(teks); i++{
    fmt.Printf("Karakter ke-%d adalah %c\n", i+1, teks[i])
  }
}
