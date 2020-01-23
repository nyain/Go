/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat fungsi untuk menukar posisi dari 3 string yang user input.
    String ke 1 ditukar menjadi string ke 3, string ke 2 ditukar menjadi string ke 1, dan string ke 3 ditukar menjadi string ke 2

*/
package main

import "fmt"

func main() {
    var (
        satu,dua,tiga string
        temp          string
    )
    fmt.Println("Nama: Vincent Williams Jonathan")
    fmt.Print("NIM: 1301190381"\n)
    fmt.Print("Masukkan input strings: ")
    fmt.Scanln(&satu)
    fmt.Print("Masukkan input strings: ")
    fmt.Scanln(&dua)
    fmt.Print("Masukkan input strings: ")
    fmt.Scanln(&tiga)
    fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
    temp = satu
    satu = dua
    dua = tiga
    tiga = temp
    fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
