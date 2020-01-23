package main

import "fmt"

var a,b,c int

func main() {
    fmt.Print("Masukkan angka pertama: ")
    fmt.Scan(&a)
    fmt.Print("Masukkan angka kedua: ")
    fmt.Scan(&b)
    c = a
    a = b
    b = c
    fmt.Println(" ")
    fmt.Print("Angka pertama: ")
    fmt.Println(a)
    fmt.Print("Angka kedua: ")
    fmt.Println(b)
}
