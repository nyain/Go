package main

import (
    "fmt"
    "math/rand"
)

func main() {
  var seed,dadang,dadu,aku int64
  fmt.Print("Masukkan angka rahasia: ")
  fmt.Scan(&seed)
  rand.Seed(seed)

  dadang = int64 (rand.Intn(6)+1)
  dadu = int64 (rand.Intn(6)+1)

  fmt.Print("Anda: ")
  fmt.Scan(&aku)
  fmt.Println("Dadang:",dadang)

  if dadu == aku {
    fmt.Println("Nilai dadu adalah",dadu,",Pemenang adalah anda")
  } else if dadu == dadang{
    fmt.Println("Nilai dadu adalah",dadu,",Pemenang adalah Dadang")
  } else if dadu == aku && dadu == dadang {
    fmt.Println("Seri")
  } else {
    fmt.Println("NIlai dadu adalah",dadu,", Tidak ada pemenang")
  }
}
