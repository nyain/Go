package main

import "fmt"

var nomor int

func main() {
	fmt.Println("MENU ")
	fmt.Println("1. Reservasi Tiket")
	fmt.Println("2. Mencari Penonton dengan Kode Reservasi")
	fmt.Println("3. Mencari Penonton dengan Nomor Identitas")
	fmt.Println("4. Data penonton di Area VIP")
	fmt.Println("5. Statistik Okupansi Area VIP,Regular Kiri,Regular Kanan")
	fmt.Println("6. Statistik Komposisi Penonton Dewasa dan Anak-Anak")
	fmt.Println("7. Menampilkan semua penonton terurut berrdasarkan Usia")
	fmt.Println("8. Jumlah kursi kosong pada area VIP,Regular Kiri dan Regular Kanan")
	fmt.Println("9. Exit")
	fmt.Print("Masukkan Nomor :")
	fmt.Scanln(&nomor)
}
