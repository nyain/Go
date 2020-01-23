/* TOPIK 4
Dibuat oleh : Kelompok 1
Hassan Rizky Putra Sailellah / 1301190328
Naufal Rezky Ananda / 1301190478

Kelas : IF-43-05
*/

package main

import "fmt"
import "math/rand"
import "os"

const N = 6
const M = 5
const K = 6

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

var umur, i, j, p, kodereservasi, anak, dewasa, noiden, jumlah int
var hargaTiketR, hargaTiketV, bayar, total float64
var nama, jenisTiket, cetak string
var s int64
var nomor int
var found, terisi bool

func inputdataidentitas() {
	fmt.Print("Nomor Identitas : ")
	fmt.Scanln(&noiden)
	fmt.Print("Nama Panggilan  : ")
	fmt.Scanln(&nama)
	fmt.Print("Umur            : ")
	fmt.Scanln(&umur)
}

func datapenonton() {

	if jenisTiket == "Regular_Kiri" || jenisTiket == "Regular_Kanan" {
		if umur >= 6 && umur < 19 {
			anak = anak + 1
			bayar = hargaTiketR * 0.6
			if i > 1 {
				bayar = bayar - ((bayar * 0.03) * float64(i))
			} else {
				bayar = bayar
			}
			terisi = true
		} else if umur >= 19 {
			dewasa = dewasa + 1
			if i > 1 {
				bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
			} else {
				bayar = hargaTiketR
			}
			terisi = true
		} else {
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
			terisi = true
		} else if umur >= 19 {
			dewasa = dewasa + 1

			if i > 1 {
				bayar = hargaTiketV - ((hargaTiketV * 0.03) * float64(i))
			} else {
				bayar = hargaTiketV
			}
			terisi = true
		} else {
			fmt.Println("Umur Belum Mencukupi")
		}
	}

	if bayar != 0 {
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

		fmt.Println("\n   THIS IS YOUR TICKET")
		fmt.Println("Kode Reservasi  :", kodereservasi)
		fmt.Println("No Tempat Duduk :", i, j)
		fmt.Println("Harga Tiket     :", bayar)
		a := i
		fmt.Print("Cetak tiket? (y/t): ")
		fmt.Scanln(&cetak)
		if cetak == "y" {
			rubahkettxt(nama, jenisTiket, kodereservasi, a, j)
		}
		fmt.Println()
		i = 0
		j = 0
		total = total + bayar
		fmt.Println()
	} else {
		fmt.Println("Ada kesalahan saat penginputan identitas")
	}
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
	if i <= N && i > 0 {
		for !terisi {
			if jenisTiket == "Regular_Kiri" {
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
			} else if jenisTiket == "Regular_Kanan" {
				j = 0
				if datakursiRR[i][j].noiden == 0 {
					inputdataidentitas()
					datapenonton()
				} else if datakursiRR[i][j].noiden != 0 {
					for j < N && !terisi {
						j = j + 1
						if datakursiRR[i][j].noiden == 0 {
							inputdataidentitas()
							datapenonton()
						} else if datakursiRR[i][4].noiden != 0 {
							kursipenuh()
						}
					}
				}
			}
		}
		p = p + 1
	} else {
		fmt.Println("Kursi yang anda pesan tidak ada")
	}
}

func outputdatapenontonVIP() {
	fmt.Println("Nomor Identitas : ", datakursiV[i][j].noiden)
	fmt.Println("Nama Panggilan  : ", datakursiV[i][j].nama)
	fmt.Println("Umur            : ", datakursiV[i][j].umur)
	fmt.Println("Kode Reservasi  : ", datakursiV[i][j].kodereservasi)
	fmt.Println("Jenis Tiket     : VIP")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}
func outputdatapenontonRegularKanan() {
	fmt.Println("Nomor Identitas : ", datakursiRR[i][j].noiden)
	fmt.Println("Nama Panggilan  : ", datakursiRR[i][j].nama)
	fmt.Println("Umur            : ", datakursiRR[i][j].umur)
	fmt.Println("Kode Reservasi  : ", datakursiRR[i][j].kodereservasi)
	fmt.Println("Jenis Tiket     : Regular Kanan")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}
func outputdatapenontonRegularKiri() {
	fmt.Println("Nomor Identitas : ", datakursiRL[i][j].noiden)
	fmt.Println("Nama Panggilan  : ", datakursiRL[i][j].nama)
	fmt.Println("Umur            : ", datakursiRL[i][j].umur)
	fmt.Println("Kode Reservasi  : ", datakursiRL[i][j].kodereservasi)
	fmt.Println("Jenis Tiket     : Regular Kiri")
	fmt.Println("Nomor Kursi     : ", i, j)
	fmt.Println()
}

func rubahkettxt(nama, jenisTiket string, kodereservasi, a, j int) {
	f, err := os.OpenFile("tiket.txt", os.O_CREATE|os.O_APPEND, 0600)
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

func main() {
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	fmt.Println("~~~~~~~APLIKASI PEMESANAN TIKET KONSER MUSIK KLASIK~~~~~~~~")
	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	for nomor != 9 {
		fmt.Println("\nMENU ")
		fmt.Println("1. Reservasi Tiket")
		fmt.Println("2. Mencari Penonton dengan Kode Reservasi")
		fmt.Println("3. Mencari Penonton dengan Nomor Identitas")
		fmt.Println("4. Data penonton di Area VIP")
		fmt.Println("5. Statistik Okupansi Area VIP,Regular Kiri,Regular Kanan")
		fmt.Println("6. Statistik Komposisi Penonton Dewasa dan Anak-Anak")
		fmt.Println("7. Menampilkan semua penonton terurut berdasarkan Usia")
		fmt.Println("8. Jumlah kursi kosong pada area VIP,Regular Kiri dan Regular Kanan")
		fmt.Println("9. Exit")
		fmt.Print("Masukkan Nomor :")
		fmt.Scanln(&nomor)

		if nomor == 1 {
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
						if i < N && j < M && i > 0 && j > 0 {
							for !terisi {
								if datakursiV[i][j].noiden == 0 {
									inputdataidentitas()
									datapenonton()
								} else if datakursiV[i][j].noiden != 0 {
									kursipenuh()
								}
							}
							p = p + 1
						} else {
							fmt.Println("Kursi yang anda pesan tidak ada")
						}
					} else if jenisTiket == "Regular_Kiri" {
						datakursiregular()

					} else if jenisTiket == "Regular_Kanan" {
						datakursiregular()
					}
				}
				fmt.Printf("Total Harga dari %v tiket : Rp. %v\n", jumlah, total)
			} else {
				fmt.Println("Melebihi Batas Maksimal")
			}
		} else if nomor == 2 {
			i = 0
			j = 0
			fmt.Println("\nMasukkan kode Reservasi :")
			fmt.Scanln(&kodereservasi)
			for i < N {
				for j < M {
					if datakursiV[i][j].kodereservasi == kodereservasi {
						outputdatapenontonVIP()
					}
					j = j + 1
				}
				j = 0
				for j < K {
					if datakursiRL[i][j].kodereservasi == kodereservasi {
						outputdatapenontonRegularKiri()
					}
					j = j + 1
				}
				j = 0
				for j < K {
					if datakursiRR[i][j].kodereservasi == kodereservasi {
						outputdatapenontonRegularKanan()
					}
					j = j + 1
				}
				j = 0
				i = i + 1
			}
		} else if nomor == 3 {
			i = 0
			j = 0
			fmt.Print("\nMasukkan Nomor Identitas :")
			fmt.Scanln(&noiden)
			for i < N {
				for j < M {
					if datakursiV[i][j].noiden == noiden {
						outputdatapenontonVIP()
					}
					j = j + 1
				}
				j = 0
				for j < K {
					if datakursiRL[i][j].noiden == noiden {
						outputdatapenontonRegularKiri()
					}
					j = j + 1
				}
				j = 0
				for j < K {
					if datakursiRR[i][j].noiden == noiden {
						outputdatapenontonRegularKanan()
					}
					j = j + 1
				}
				j = 0
				i = i + 1
			}
		} else if nomor == 4 {
			fmt.Print("\nMasukkan Baris Kursi :")
			fmt.Scanln(&i)
			fmt.Print("Masukkan Nomor Kursi :")
			fmt.Scanln(&j)

			if datakursiV[i][j].kodereservasi != 0 {
				outputdatapenontonVIP()
			} else {
				fmt.Println("Kursi Tersebut KOSONG.")
			}
		} else if nomor == 5 {
			vip := 0
			kanan := 0
			kiri := 0
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					if datakursiV[i][j].kodereservasi != 0 {
						vip++
					}
				}
			}
			for i := 0; i < N; i++ {
				for j := 0; j < K; j++ {
					if datakursiRR[i][j].kodereservasi != 0 {
						kanan++
					}
					if datakursiRL[i][j].kodereservasi != 0 {
						kiri++
					}
				}
			}
			fmt.Println("Isi Kursi VIP: ", vip)
			fmt.Println("Isi Kursi Kanan: ", kanan)
			fmt.Println("Isi Kursi Kiri: ", kiri)

		} else if nomor == 6 {
			fmt.Println("Jumlah Penonton Anak-anak: ", anak)
			fmt.Println("Jumlah Penonton Dewasa: ", dewasa)

		} else if nomor == 7 {
			datakursibaruV = datakursiV
			datakursibaruRL = datakursiRL
			datakursibaruRR = datakursiRR

			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					min := i
					min2 := j
					for k := 0; k < N; k++ {
						for l := 0; l < M; l++ {
							if datakursibaruV[min][min2].umur < datakursibaruV[k][l].umur {
								min = k
								min2 = l
							}
							temp := datakursibaruV[i][j]
							datakursibaruV[i][j] = datakursibaruV[min][min2]
							datakursibaruV[min][min2] = temp
						}
					}
				}
			}

			for i := 0; i < N; i++ {
				for j := 0; j < K; j++ {
					min := i
					min2 := j
					for k := 0; k < N; k++ {
						for l := 0; l < K; l++ {
							if datakursibaruRR[min][min2].umur < datakursibaruRR[k][l].umur {
								min = k
								min2 = l
							}
							temp := datakursibaruRR[i][j]
							datakursibaruRR[i][j] = datakursibaruRR[min][min2]
							datakursibaruRR[min][min2] = temp

							if datakursibaruRL[min][min2].umur < datakursibaruRL[k][l].umur {
								min = k
								min2 = l
							}
							temp1 := datakursibaruRL[i][j]
							datakursibaruRL[i][j] = datakursibaruRL[min][min2]
							datakursibaruRL[min][min2] = temp1
						}
					}
				}
			}

			fmt.Println("VIP:")
			for i := 0; i < N; i++ {
				for j := 0; j < M; j++ {
					if datakursibaruV[i][j].kodereservasi != 0 {
						fmt.Println(datakursibaruV[i][j])
					}
				}
			}
			fmt.Println("Kursi Kanan: ")
			for i := 0; i < N; i++ {
				for j := 0; j < K; j++ {
					if datakursibaruRR[i][j].kodereservasi != 0 {
						fmt.Println(datakursibaruRR[i][j])
					}
				}
			}
			fmt.Println("Kursi Kiri: ")
			for i := 0; i < N; i++ {
				for j := 0; j < K; j++ {
					if datakursibaruRL[i][j].kodereservasi != 0 {
						fmt.Println(datakursibaruRL[i][j])
					}
				}
			}

		} else if nomor == 8 {
			vip := 0
			kanan := 0
			kiri := 0
			for i := 1; i < N; i++ {
				for j := 1; j < M; j++ {
					if datakursiV[i][j].kodereservasi == 0 {
						vip++
					}
				}
				j = 1

				for j := 1; j < K; j++ {
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

		}
		fmt.Println("\nPress Enter to Continue.... ")
		fmt.Scanln()
	}
	fmt.Println("Created By: ")
	fmt.Println("Naufal Rezky Ananda / 1301190478")
	fmt.Println("Hassan Rizky Putra Sailellah / 1301190328")
	fmt.Println("\nKelas : IF-43-05")
	fmt.Println("\nThanks You!!!")
}
