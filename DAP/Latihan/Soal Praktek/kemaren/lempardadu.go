package main

import (
    "fmt"
    "math/rand"
)

func main() {
  var (
    seed,dadu int64
    aku,nilai string
  )
  fmt.Print("Masukkan angka rahasia: ")
  fmt.Scan(&seed)
  rand.Seed(seed)

  dadu = int64 (rand.Intn(6)+1)

  fmt.Print("Anda: ")
  fmt.Scan(&aku,&nilai)

  if aku == "genap" && nilai == "kecil" && dadu % 2 == 0 && dadu >= 2 && dadu < 4  {
    fmt.Println("Nilai dadu adalah",dadu,"Anda menang")
  } else if aku == "genap" && nilai == "besar" && dadu % 2 == 0 && dadu >=4 && dadu <=6 {
    fmt.Println("Nilai dadu adalah",dadu,"Anda menang")
  } else if aku == "ganjil" && nilai == "kecil" && dadu % 2 != 0 && dadu >= 1 && dadu <= 3   {
    fmt.Println("Nilai dadu adalah",dadu,"Anda menang")
  } else if aku == "ganjil" && nilai == "besar" && dadu % 2 != 0 && dadu > 3 && dadu <= 5{
    fmt.Println("Nilai dadu adalah",dadu,"Anda menang")
  } else {
      fmt.Println("Nilai dadu adalah",dadu,"Anda kalah")
  }
}
