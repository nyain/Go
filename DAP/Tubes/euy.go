package main

import "fmt"
import "math/rand"
import "os"

const N = 5
const M = 4
const K = 5

type data struct {
	nama                string
	noiden              int
	kodereservasi, umur int
}

var datakursiV [N][M]data
var datakursiRL [N][K]data
var datakursiRR [N][K]data
var datakursibaruV [N][M]data
var datakursibaruRR [N][K]data
var datakursibaruRL [N][K]data

var umur, i, j, p, kodereservasi, anak, dewasa, noiden, jumlah, nomor int
var hargaTiketR, hargaTiketV, bayar, total float64
var nama, jenisTiket, cetak string
var s int64
var found, terisi bool

func bngst(nama, jenisTiket string, kodereservasi, a, j int) {
	f, err := os.OpenFile("bngst.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	for i := 0; i < 1; i++ {
		_, err = f.WriteString(fmt.Sprintf("Nama: %s\nKelas: %s\nKode Reservasi: %d\nTempat Duduk: %d %d\n\n", nama, jenisTiket, kodereservasi, a, j))
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}

}
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
	a := i
	fmt.Print("Cetak tiket? (y/t): ")
	fmt.Scan(&cetak)
	if cetak == "y" {
		bngst(nama, jenisTiket, kodereservasi, a, j)
	}
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
		if datakursiRL[i][j].noiden == 0 {
			inputdataidentitas()
			datapenonton()
		} else if datakursiRL[i][j].noiden != 0 {
			for j < N && !terisi {
				j = j + 1
				if datakursiRL[i][j].noiden == 0 {
					inputdataidentitas()
					datapenonton()
				} else if datakursiRL[i][4].noiden != 0 {
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
			if jenisTiket == "VIP" || jenisTiket == "Regular_Kiri" || jenisTiket == "Regular_Kanan" {
				if jenisTiket == "VIP" {
					fmt.Println("Pilih Tempat Duduk")
					fmt.Print("Baris Kursi ke : ")
					fmt.Scanln(&i)
					fmt.Print("Nomor Kursi ke : ")
					fmt.Scanln(&j)
					for !terisi {
						if datakursiV[i][j].noiden == 0 {
							inputdataidentitas()
							datapenonton()
						} else if datakursiV[i][j].noiden != 0 {
							kursipenuh()
						}
					}
				} else if jenisTiket == "Regular_Kiri" {
					datakursiregular()

				} else if jenisTiket == "Regular_Kanan" {
					datakursiregular()
				}
				p = p + 1
			} else {
				fmt.Println("Data salah")
			}
		}
	} else {
		fmt.Println("Melebihi Batas Maksimal")
	}
	fmt.Printf("Total Harga dari %v tiket : Rp. %v\n", jumlah, total)
	menu()
}
func searchreserve() {
	i = 0
	j = 0
	fmt.Print("\nMasukkan kode Reservasi :")
	fmt.Scanln(&kodereservasi)
	for i < N {
		for j < M {
			if datakursiV[i][j].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiV[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiV[i][j].nama)
				fmt.Println("Umur            : ", datakursiV[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiV[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : VIP")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRL[i][j].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiRL[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRL[i][j].nama)
				fmt.Println("Umur            : ", datakursiRL[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRL[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kiri")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRR[i][j].kodereservasi == kodereservasi {
				fmt.Println("Nomor Identitas : ", datakursiRR[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRR[i][j].nama)
				fmt.Println("Umur            : ", datakursiRR[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRR[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kanan")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
	menu()
}
func searchidentity() {

	i = 0
	j = 0
	fmt.Print("\nMasukkan Nomor Identitas :")
	fmt.Scanln(&noiden)
	for i < N {
		for j < M {
			if datakursiV[i][j].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiV[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiV[i][j].nama)
				fmt.Println("Umur            : ", datakursiV[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiV[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : VIP")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRL[i][j].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiRL[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRL[i][j].nama)
				fmt.Println("Umur            : ", datakursiRL[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRL[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kiri")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if datakursiRR[i][j].noiden == noiden {
				fmt.Println("Nomor Identitas : ", datakursiRR[i][j].noiden)
				fmt.Println("Nama Panggilan  : ", datakursiRR[i][j].nama)
				fmt.Println("Umur            : ", datakursiRR[i][j].umur)
				fmt.Println("Kode Reservasi  : ", datakursiRR[i][j].kodereservasi)
				fmt.Println("Jenis Tiket     : Regular Kanan")
				fmt.Println("Nomor Kursi     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
	menu()
}
func datavip() {
	fmt.Print("\nMasukkan Baris Kursi :")
	fmt.Scanln(&i)
	fmt.Print("Masukkan Nomor Kursi :")
	fmt.Scanln(&j)

	fmt.Println("\nNomor Identitas : ", datakursiV[i][j].noiden)
	fmt.Println("Nama Panggilan  : ", datakursiV[i][j].nama)
	fmt.Println("Umur            : ", datakursiV[i][j].umur)
	fmt.Println("Kode Reservasi  : ", datakursiV[i][j].kodereservasi)
	fmt.Println("Jenis Tiket     : VIP")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println("")
	menu()
}
func okupansi() {
	vip := 0
	kanan := 0
	kiri := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if datakursiV[i][j].kodereservasi > 0 {
				vip++
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			if datakursiRR[i][j].kodereservasi > 0 {
				kanan++
			}
			if datakursiRL[i][j].kodereservasi > 0 {
				kiri++
			}
		}
	}
	fmt.Println("Isi Kursi VIP: ", vip)
	fmt.Println("Isi Kursi Kanan: ", kanan)
	fmt.Println("Isi Kursi Kiri: ", kiri)
	menu()
}
func komposisi() {
	fmt.Println("Jumlah Penonton Anak-anak: ", anak)
	fmt.Println("Jumlah Penonton Dewasa: ", dewasa)
	menu()
}
func sortusia() {
	datakursibaruV = datakursiV
	datakursibaruRL = datakursiRL
	datakursibaruRR = datakursiRR

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			cur := i
			cur2 := j
			for k := 0; k < N; k++ {
				for l := 0; l < M; l++ {
					if datakursibaruV[cur][cur2].umur < datakursibaruV[k][l].umur {
						cur = k
						cur2 = l
					}
					temp := datakursibaruV[i][j]
					datakursibaruV[i][j] = datakursibaruV[cur][cur2]
					datakursibaruV[cur][cur2] = temp
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			cur := i
			cur2 := j
			for k := 0; k < N; k++ {
				for l := 0; l < K; l++ {
					if datakursibaruRR[cur][cur2].umur < datakursibaruRR[k][l].umur {
						cur = k
						cur2 = l
					}
					temp := datakursibaruRR[i][j]
					datakursibaruRR[i][j] = datakursibaruRR[cur][cur2]
					datakursibaruRR[cur][cur2] = temp
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			cur := i
			cur2 := j
			for k := 0; k < N; k++ {
				for l := 0; l < K; l++ {
					if datakursibaruRL[cur][cur2].umur < datakursibaruRL[k][l].umur {
						cur = k
						cur2 = l
					}
					temp := datakursibaruRL[i][j]
					datakursibaruRL[i][j] = datakursibaruRL[cur][cur2]
					datakursibaruRL[cur][cur2] = temp
				}
			}
		}
	}

	fmt.Println("VIP:")
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Println(datakursibaruV[i][j])
			fmt.Println(datakursibaruV[i][j])
		}
	}
	fmt.Println("Kursi Kanan: ")
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Println(datakursibaruRR[i][j])
		}
	}
	fmt.Println("Kursi Kiri: ")
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Println(datakursibaruRL[i][j])
		}
	}
	menu()
}
func empty() {
	vip := 0
	kanan := 0
	kiri := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if datakursiV[i][j].kodereservasi == 0 {
				vip++
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			if datakursiRR[i][j].kodereservasi == 0 {
				kanan++
			}
			if datakursiRL[i][j].kodereservasi == 0 {
				kiri++
			}
		}
	}
	fmt.Println("Jumlah Kursi Kosong di VIP: ", vip)
	fmt.Println("Jumlah Kursi Kosong di Kanan: ", kanan)
	fmt.Println("Jumlah Kursi Kosong di Kiri: ", kiri)
	menu()
}

func main() {
	menu()
	for nomor >= 1 && nomor <= 8 {
		switch {
		case nomor == 1:
			reserve()
		case nomor == 2:
			searchreserve()
		case nomor == 3:
			searchidentity()
		case nomor == 4:
			datavip()
		case nomor == 5:
			okupansi()
		case nomor == 6:
			komposisi()
		case nomor == 7:
			sortusia()
		case nomor == 8:
			empty()
		}
	}
}
