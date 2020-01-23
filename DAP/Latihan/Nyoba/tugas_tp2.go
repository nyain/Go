/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat fungsi untuk mencari tahu sebuah tahun kabisat atau bukan.
    Tahun kabisat adalah tahun yang habis dibagi 400 atau 4 tetapi tidak habis dibagi 100.

*/

package main

import "fmt"

func main() {
    var (
        tahun int
        pengecekan bool
    )
    fmt.Println("Nama: Vincent Williams Jonathan")
    fmt.Println("NIM: 1301190381")

    fmt.Print("Masukkan tahun: ")
    fmt.Scanln(&tahun)

    pengecekan = tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 != 0

    fmt.Print("Kabisat: ")
    fmt.Println(pengecekan)
}
