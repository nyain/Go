/*
				KELOMPOK 5
				 IF-43-05
	 NAUFAL YOZHA REVANSYAH
	 			1301194282
	VINCENT WILLIAMS JONATHAN
				1301190381
*/

package main

import "fmt"
import "math/rand"
import "os"

const N = 8 //dikurang 1
const M = 5 //dikurang 1
const K = 8 //dikurang 1

type tiket struct {
	nama                        string
	identity, reservecode, usia int
}

var VIP [N][M]tiket
var RegKiri [N][K]tiket
var RegKanan [N][K]tiket

var tampungVIP [N][M]tiket
var tampungRegKiri [N][K]tiket
var tampungRegKanan [N][K]tiket

var usia, i, j, p, reservecode, anak, dewasa, identity, jumlah int
var hargaR, hargaVIP, bayar, total float64
var nama, jenisTiket string
var seed int64
var nomor int
var found, full bool

func menu() {
	fmt.Println("SELAMAT DATANG DI KONSER PACHELBEL")
	fmt.Println("1. Pesan tiket")
	fmt.Println("2. Mencari data penonton dengan kode reservasi")
	fmt.Println("3. Mencari data penonton dengan nomor identitas")
	fmt.Println("4. Data penonton di area VIP")
	fmt.Println("5. Statistik okupansi dan komposisi di semua area")
	fmt.Println("6. Menampilkan semua penonton terurut berdasarkan usia")
	fmt.Println("7. Jumlah kursi kosong pada semua area")
	fmt.Println("8. Keluar")
	fmt.Print("Masukkan Nomor :")
	fmt.Scan(&nomor)
}
func inputdata() {
	fmt.Print("Nomor Identitas : ")
	fmt.Scan(&identity)
	fmt.Print("Nama            : ")
	fmt.Scan(&nama)
	fmt.Print("Usia            : ")
	fmt.Scan(&usia)
}
func data() {
	if jenisTiket == "RegKiri" || jenisTiket == "RegKanan" {
		if usia >= 6 && usia < 19 {
			anak++
			bayar = hargaR * 0.6
			if i > 1 {
				bayar = bayar - ((bayar * 0.03) * float64(i))
			} else {
				bayar = bayar
			}
		} else if usia >= 19 {
			dewasa++
			if i > 1 {
				bayar = hargaR - ((hargaR * 0.03) * float64(i))
			} else {
				bayar = hargaR
			}
		} else if usia > 0 {
			fmt.Println("Usia belum mencukupi")
		}
	} else {
		if usia >= 6 && usia < 19 {
			anak++
			bayar = hargaVIP * 0.6
			if i > 1 {
				bayar = bayar - ((bayar * 0.03) * float64(i))
			} else {
				bayar = bayar
			}
		} else if usia >= 19 {
			dewasa++

			if i > 1 {
				bayar = hargaVIP - ((hargaVIP * 0.03) * float64(i))
			} else {
				bayar = hargaVIP
			}
		} else if usia > 0 {
			fmt.Println("Usia belum mencukupi")
		}
		if bayar != 0 {
			if jenisTiket == "VIP" {
				VIP[i][j].identity = identity
				VIP[i][j].nama = nama
				VIP[i][j].usia = usia
				VIP[i][j].reservecode = reservecode
			} else if jenisTiket == "RegKiri" {
				RegKiri[i][j].identity = identity
				RegKiri[i][j].nama = nama
				RegKiri[i][j].usia = usia
				RegKiri[i][j].reservecode = reservecode
			} else if jenisTiket == "RegKanan" {
				RegKanan[i][j].identity = identity
				RegKanan[i][j].nama = nama
				RegKanan[i][j].usia = usia
				RegKanan[i][j].reservecode = reservecode
			}
			fmt.Println("\n   INVOICE")
			fmt.Println("Kode Reservasi  :", reservecode)
			fmt.Println("No Tempat Duduk :", i, j)
			fmt.Println("Harga Tiket     :", bayar)
			a := i
			cetakTiket(nama, jenisTiket, reservecode, a, j)
			fmt.Println()
			i = 0
			j = 0
			full = true
			total = total + bayar
			fmt.Println()
		} else {
			fmt.Println("SILAHKAN INPUT ULANG")
		}
	}
}
func inputkursireguler() {
	fmt.Println("Pilih Tempat Duduk")
	fmt.Print("Baris Kursi ke : ")
	fmt.Scan(&i)
	if i <= N && i > 0 {
		for !full {
			if jenisTiket == "RegKiri" {
				j = 0
				if RegKiri[i][j].identity == 0 {
					inputdata()
					data()
				} else if RegKiri[i][j].identity != 0 {
					for j < N && !full {
						j = j + 1
						if RegKiri[i][j].identity == 0 {
							inputdata()
							data()
						} else if RegKiri[i][7].identity != 0 {
							overload()
						}
					}
				}
			} else if jenisTiket == "RegKanan" {
				j = 0
				if RegKanan[i][j].identity == 0 {
					inputdata()
					data()
				} else if RegKanan[i][j].identity != 0 {
					for j < N && !full {
						j = j + 1
						if RegKanan[i][j].identity == 0 {
							inputdata()
							data()
						} else if RegKanan[i][7].identity != 0 {
							overload()
						}
					}
				}
			}
		}
		p++
	} else {
		fmt.Println("Kursi yang anda pesan tidak ada")
	}
}
func dataVIP() {
	fmt.Println("\nNomor Identitas : ", VIP[i][j].identity)
	fmt.Println("Nama     	: ", VIP[i][j].nama)
	fmt.Println("Usia            : ", VIP[i][j].usia)
	fmt.Println("Kode Reservasi  : ", VIP[i][j].reservecode)
	fmt.Println("Jenis Tiket     :  VIP")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}
func dataRegKiri() {
	fmt.Println("\nNomor Identitas : ", RegKiri[i][j].identity)
	fmt.Println("Nama     	: ", RegKiri[i][j].nama)
	fmt.Println("Usia            : ", RegKiri[i][j].usia)
	fmt.Println("Kode Reservasi  : ", RegKiri[i][j].reservecode)
	fmt.Println("Jenis Tiket     :  Regular Kiri")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}
func dataRegKanan() {
	fmt.Println("\nNomor Identitas : ", RegKanan[i][j].identity)
	fmt.Println("Nama     	: ", RegKanan[i][j].nama)
	fmt.Println("Usia            : ", RegKanan[i][j].usia)
	fmt.Println("Kode Reservasi  : ", RegKanan[i][j].reservecode)
	fmt.Println("Jenis Tiket     :  Regular Kanan")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}
func overload() {
	if jenisTiket == "VIP" {
		fmt.Print("PENUH,INPUT ULANG\n\n")
		fmt.Print("Baris Kursi ke : ")
		fmt.Scan(&i)
		fmt.Print("Nomor Kursi ke : ")
		fmt.Scan(&j)
	} else {
		fmt.Print("PENUH,INPUT ULANG\n\n")
		fmt.Print("Baris Kursi ke : ")
		fmt.Scan(&i)
	}
}
func cetakTiket(nama, jenisTiket string, reservecode, a, j int) {
	f, err := os.OpenFile("Ticket.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer f.Close()
	_, err = f.WriteString(fmt.Sprintf("Nama: %s\nKelas: %s\nKode Reservasi: %d\nTempat Duduk: %d %d\n\n", nama, jenisTiket, reservecode, a, j))
	if err != nil {
		fmt.Printf("error writing string: %v", err)
	}
}

func reserve() {
	fmt.Print("Jumlah tiket yang ingin dipesan (maksimal 10 tiket) :")
	fmt.Scan(&jumlah)
	if jumlah <= 10 {
		p = 0
		total = 0
		seed++
		rand.Seed(seed)
		reservecode = rand.Intn(1000000)
		for p < jumlah {
			full = false
			bayar = 0
			hargaVIP = 300000
			hargaR = 100000
			fmt.Print("Jenis Tiket(VIP/RegKiri/RegKanan) :")
			fmt.Scan(&jenisTiket)
			if jenisTiket == "VIP" {
				fmt.Print("Baris Kursi ke : ")
				fmt.Scan(&i)
				fmt.Print("Nomor Kursi ke : ")
				fmt.Scan(&j)
				if i < N && j < M && i > 0 && j > 0 {
					for !full {
						if VIP[i][j].identity == 0 {
							inputdata()
							data()
						} else if VIP[i][j].identity != 0 {
							overload()
						}
					}
					p++
				} else {
					fmt.Println("KURSI TIDAK ADA, SILAHKAN INPUT ULANG")
				}
			} else if jenisTiket == "RegKiri" {
				inputkursireguler()

			} else if jenisTiket == "RegKanan" {
				inputkursireguler()
			}
		}
		fmt.Printf("Total Harga dari %v tiket : Rp. %v\n", jumlah, total)
	} else {
		fmt.Println("Melebihi Batas Maksimal")
	}
	main()
}
func searchreserve() {
	i = 0
	j = 0
	fmt.Print("\nMasukkan kode Reservasi :")
	fmt.Scan(&reservecode)
	for i < N {
		for j < M {
			if VIP[i][j].reservecode == reservecode {
				dataVIP()
			}
			j++
		}
		j = 0
		for j < K {
			if RegKiri[i][j].reservecode == reservecode {
				dataRegKiri()
			}
			j++
		}
		j = 0
		for j < K {
			if RegKanan[i][j].reservecode == reservecode {
				dataRegKanan()
			}
			j++
		}
		j = 0
		i++
	}
	main()
}
func searchidentity() {
	i = 0
	j = 0
	fmt.Print("\nMasukkan Nomor Identitas :")
	fmt.Scan(&identity)
	for i < N {
		for j < M {
			if VIP[i][j].identity == identity {
				dataVIP()
			}
			j++
		}
		j = 0
		for j < K {
			if RegKiri[i][j].identity == identity {
				dataRegKiri()
			}
			j++
		}
		j = 0
		for j < K {
			if RegKanan[i][j].identity == identity {
				dataRegKanan()
			}
			j++
		}
		j = 0
		i++
	}
	main()
}
func datavip() {
	fmt.Print("\nMasukkan Baris Kursi :")
	fmt.Scan(&i)
	fmt.Print("Masukkan Nomor Kursi :")
	fmt.Scan(&j)

	if VIP[i][j].usia != 0 {
		dataVIP()
	} else {
		fmt.Println("\nKursi Tersebut KOSONG.")
	}
	main()
}
func okupansi() {
	vip := 0
	kanan := 0
	kiri := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if VIP[i][j].reservecode != 0 {
				vip++
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			if RegKanan[i][j].reservecode != 0 {
				kanan++
			}
			if RegKiri[i][j].reservecode != 0 {
				kiri++
			}
		}
	}
	fmt.Println("\nIsi Kursi VIP: ", vip)
	fmt.Println("Isi Kursi Kiri: ", kiri)
	fmt.Println("Isi Kursi Kanan: ", kanan)
	fmt.Println("Jumlah Penonton Anak-anak: ", anak)
	fmt.Println("Jumlah Penonton Dewasa: ", dewasa, "\n")
	main()
}
func sortusia() {
	tampungVIP = VIP
	tampungRegKiri = RegKiri
	tampungRegKanan = RegKanan

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			x := i
			y := j
			for k := 0; k < N; k++ {
				for l := 0; l < M; l++ {
					if tampungVIP[x][y].usia < tampungVIP[k][l].usia {
						x = k
						y = l
					}
					temp := tampungVIP[i][j]
					tampungVIP[i][j] = tampungVIP[x][y]
					tampungVIP[x][y] = temp
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			x := i
			y := j
			for k := 0; k < N; k++ {
				for l := 0; l < K; l++ {
					if tampungRegKanan[x][y].usia < tampungRegKanan[k][l].usia {
						x = k
						y = l
					}
					temp := tampungRegKanan[i][j]
					tampungRegKanan[i][j] = tampungRegKanan[x][y]
					tampungRegKanan[x][y] = temp
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			x := i
			y := j
			for k := 0; k < N; k++ {
				for l := 0; l < K; l++ {
					if tampungRegKiri[x][y].usia < tampungRegKiri[k][l].usia {
						x = k
						y = l
					}
					temp := tampungRegKiri[i][j]
					tampungRegKiri[i][j] = tampungRegKiri[x][y]
					tampungRegKiri[x][y] = temp
				}
			}
		}
	}
	fmt.Println("\nVIP:")
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if tampungVIP[i][j].reservecode != 0 {
				fmt.Println(tampungVIP[i][j])
			}
		}
	}
	fmt.Println("\nKursi Kiri: ")
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			if tampungRegKiri[i][j].reservecode != 0 {
				fmt.Println(tampungRegKiri[i][j])
			}
		}
	}
	fmt.Println("\nKursi Kanan: ")
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			if tampungRegKanan[i][j].reservecode != 0 {
				fmt.Println(tampungRegKanan[i][j])
			}
		}
	}
	main()
}
func empty() {
	vip := 0
	kanan := 0
	kiri := 0
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if VIP[i][j].identity == 0 {
				vip++
			}
		}
		j = 1

		for j := 1; j < K; j++ {
			if RegKiri[i][j].identity == 0 {
				kiri++
			}
			if RegKanan[i][j].identity == 0 {
				kanan++
			}
		}
	}

	fmt.Println("\nJumlah kursi kosong di area VIP: ", vip)
	fmt.Println("Jumlah Kursi Kosong di area reguler kiri: ", kiri)
	fmt.Println("Jumlah Kursi kosong di area reguler kanan: ", kanan, "\n")
	main()
}

func main() {
	menu()
	for nomor >= 1 && nomor < 8 {
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
			sortusia()
		case nomor == 7:
			empty()
		}
	}
}
