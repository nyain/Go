package main

import (
  "bufio"
  "os"
  "fmt"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)

  fmt.Print("Masukkan nama: ")
  scanner.Scan()

  fmt.Println(scanner.Text())
}
