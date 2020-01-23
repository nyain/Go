package main

import "fmt"
import "math/rand"

const N = 5
const M = 4
const K = 5

type data struct {
	nama                string
	noiden              int
	kodereservasi, umur int
}

var datakursiV [M][N]data
var datakursiRL [K][N]data
var datakursiRR [K][N]data

var umur, i, j, p, kodereservasi, anak, dewasa, noiden, jumlah, nomor int
var hargaTiketR, hargaTiketV, bayar, total float64
var nama, jenisTiket string
var s int64
var found, terisi bool

func menu() {
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
	fmt.Scan(&nomor)
}
func inputdataidentitas() {
	fmt.Print("Nomor Identitas : ")
	fmt.Scanln(&noiden)
	fmt.Print("Nama Panggilan  : ")
	fmt.Scanln(&nama)
	fmt.Print("Umur            : ")
	fmt.Scanln(&umur)
}
func datapenonton() {
	if jenisTiket == "VIP" {
		datakursiV[i][j].noiden = noiden
		datakursiV[i][j].nama = nama
		datakursiV[i][j].umur = umur
		datakursiV[i][j].kodereservasi = kodereservasi
	} else if jenisTiket == "Regular_Kiri" {
		datakursiRL[i][j].noiden = noiden
		datakursiRL[i][j].nama = nama
		datakursiRL[i][j].umur = umur
		datakursiRL[i][j].kodereservasi = kodereservasi
	} else if jenisTiket == "Regular_Kanan" {
		datakursiRR[i][j].noiden = noiden
		datakursiRR[i][j].nama = nama
		datakursiRR[i][j].umur = umur
		datakursiRR[i][j].kodereservasi = kodereservasi
	}

	if jenisTiket == "Regular_Kiri" || jenisTiket == "Regular_Kanan" {
		if umur >= 6 && umur < 19 {
			anak = anak + 1
			bayar = hargaTiketR * 0.6
			if i > 1 {
				bayar = bayar - ((bayar * 0.03) * float64(i))
			} else {
				bayar = bayar
			}
		} else if umur >= 19 {
			dewasa = dewasa + 1
			if i > 1 {
				bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
			} else {
				bayar = hargaTiketR
			}
		} else if umur > 0 {
			fmt.Println("Umur Belum Mencukupi")
		}
	} else {
		if umur >= 6 && umur < 19 {
			anak = anak + 1

			bayar = hargaTiketV * 0.6
			if i > 1 {
				bayar = bayar - ((bayar * 0.03) * float64(i))
			} else {
				bayar = bayar
			}
		} else if umur >= 19 {
			dewasa = dewasa + 1

			if i > 1 {
				bayar = hargaTiketV - ((hargaTiketV * 0.03) * float64(i))
			} else {
				bayar = hargaTiketV
			}
		} else if umur > 0 {
			fmt.Println("Umur Belum Mencukupi")
		}
	}

	fmt.Println("\n   THIS IS YOUR TICKET")
	fmt.Println("Kode Reservasi  :", kodereservasi)
	fmt.Println("No Tempat Duduk :", i, j)
	fmt.Println("Harga Tiket     :", bayar)
	fmt.Println()
	i = 0
	j = 0
	terisi = true
	total = total + bayar
	fmt.Println()
}
func kursipenuh() {
	if jenisTiket == "VIP" {
		fmt.Print("PENUH\n\n")
		fmt.Print("Baris Kursi ke : ")
		fmt.Scanln(&i)
		fmt.Print("Nomor Kursi ke : ")
		fmt.Scanln(&j)
	} else {
		fmt.Print("PENUH\n\n")
		fmt.Print("Baris Kursi ke : ")
		fmt.Scanln(&i)
	}
}
func datakursiregular() {
	fmt.Println("Pilih Tempat Duduk")
	fmt.Print("Baris Kursi ke : ")
	fmt.Scanln(&i)
	for !terisi {
		j = 0
		if datakursiRL[j][i].noiden == 0 {
			inputdataidentitas()
			datapenonton()
		} else if datakursiRL[j][i].noiden != 0 {
			for j < N && !terisi {
				j = j + 1
				if datakursiRL[j][i].noiden == 0 {
					inputdataidentitas()
					datapenonton()
				} else if datakursiRL[4][i].noiden != 0 {
					kursipenuh()
				}
			}
		}
	}
}
func reserve() {

	fmt.Print("Masukkan Jumlah Tiket yang dipesan(Maks.4) :")
	fmt.Scanln(&jumlah)
	if jumlah <= 4 {
		p = 0
		total = 0
		s = s + 1
		rand.Seed(s)
		kodereservasi = rand.Intn(1000000)
		for p < jumlah {
			terisi = false
			bayar = 0
			hargaTiketV = 150000
			hargaTiketR = 100000
			fmt.Print("Jenis Tiket(VIP/Regular_Kiri/Regular_Kanan) :")
			fmt.Scanln(&jenisTiket)
			if jenisTiket == "VIP" {
				fmt.Println("Pilih Tempat Duduk")
				fmt.Print("Baris Kursi ke : ")
				fmt.Scanln(&i)
				fmt.Print("Nomor Kursi ke : ")
				fmt.Scanln(&j)
				for !terisi {
					if datakursiV[j][i].noiden == 0 {
						inputdataidentitas()
						datapenonton()
					} else if datakursiV[j][i].noiden != 0 {
						kursipenuh()
					}
				}
			} else if jenisTiket == "Regular_Kiri" {
				datakursiregular()

			} else if jenisTiket == "Regular_Kanan" {
				datakursiregular()
			}
			p = p + 1
		}
		fmt.Printf("Total Harga dari %v tiket : Rp. %v\n", jumlah, total)
	} else {
		fmt.Println("Melebihi Batas Maksimal")
	}
}
func searchreserve() {
	i = 0
	j = 0
	fmt.Print("\nMasukkan kode Reservasi :")
	fmt.Scanln(&kodereservasi)
	for i < N {
		for j < M {
			if datakursiV[j][i].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiV[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiV[j][i].nama)
				fmt.Println("Umur            : ", datakursiV[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiV[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : VIP")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRL[j][i].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiRL[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRL[j][i].nama)
				fmt.Println("Umur            : ", datakursiRL[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRL[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kiri")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRR[j][i].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiRR[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRR[j][i].nama)
				fmt.Println("Umur            : ", datakursiRR[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRR[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kanan")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
}
func searchidentity() {

	i = 0
	j = 0
	fmt.Print("\nMasukkan Nomor Identitas :")
	fmt.Scanln(&noiden)
	for i < N {
		for j < M {
			if datakursiV[j][i].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiV[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiV[j][i].nama)
				fmt.Println("Umur            : ", datakursiV[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiV[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : VIP")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRL[j][i].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiRL[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRL[j][i].nama)
				fmt.Println("Umur            : ", datakursiRL[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRL[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kiri")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRR[j][i].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiRR[j][i].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRR[j][i].nama)
				fmt.Println("Umur            : ", datakursiRR[j][i].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRR[j][i].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kanan")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
}
func datavip() {
	fmt.Print("\nMasukkan Baris Kursi :")
	fmt.Scanln(&i)
	fmt.Print("Masukkan Nomor Kursi :")
	fmt.Scanln(&j)

	fmt.Println("\nNomor Identitas : ", datakursiV[j][i].noiden)
	fmt.Println("Nama Panggilan  : ", datakursiV[j][i].nama)
	fmt.Println("Umur            : ", datakursiV[j][i].umur)
	fmt.Println("Kode Reservasi  : ", datakursiV[j][i].kodereservasi)
	fmt.Println("Jenis Tiket     : VIP")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println("")
}

func main() {
	menu()
	switch {
	case nomor == 1:
		reserve()
	case nomor == 2:
		searchreserve()
	case nomor == 3:
		searchidentity()
	case nomor == 4:
		datavip()
	}
}
