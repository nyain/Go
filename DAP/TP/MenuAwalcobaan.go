package main

import "fmt"
import "math/rand"

const N = 6
const M = 5
const K = 6

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
var found bool

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
func reserve() {
	fmt.Print("Masukkan Jumlah Tiket yang dipesan(Maks.4) :")
	fmt.Scanln(&jumlah)
	p = 0
	total = 0
	s = s + 1
	rand.Seed(s)
	kodereservasi = rand.Intn(1000000)
	for p < jumlah {

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
			for i > 0 && j > 0 {
				if datakursiV[i][j].noiden == 0 {
					fmt.Print("Nomor Identitas : ")
					fmt.Scanln(&noiden)
					fmt.Print("Nama Panggilan  : ")
					fmt.Scanln(&nama)
					fmt.Print("Umur            : ")
					fmt.Scanln(&umur)
					if umur >= 6 && umur < 19 {
						anak = anak + 1
						datakursiV[i][j].noiden = noiden
						datakursiV[i][j].nama = nama
						datakursiV[i][j].umur = umur

						bayar = hargaTiketV * 0.6
						if i > 1 {
							bayar = bayar - ((bayar * 0.03) * float64(i))
						} else {
							bayar = bayar
						}

						datakursiV[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi  :", kodereservasi)
						fmt.Println("No Tempat Duduk :", i, j)
						fmt.Println("Harga Tiket     :", bayar)
						total = total + bayar
						fmt.Println()
						i = 0
						j = 0
					} else if umur >= 19 {
						dewasa = dewasa + 1
						datakursiV[i][j].noiden = noiden
						datakursiV[i][j].nama = nama
						datakursiV[i][j].umur = umur
						if i > 1 {
							bayar = hargaTiketV - ((hargaTiketV * 0.03) * float64(i))
						} else {
							bayar = hargaTiketV
						}

						datakursiV[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi  :", kodereservasi)
						fmt.Println("No Tempat Duduk :", i, j)
						fmt.Println("Harga Tiket     :", bayar)
						total = total + bayar
						fmt.Println()

						i = 0
						j = 0
					} else if umur > 0 {
						fmt.Println("Umur Belum Mencukupi")
						i = 0
						j = 0
					}
				} else if datakursiV[i][j].noiden != 0 {
					fmt.Print("PENUH\n\n")
					fmt.Print("Baris Kursi ke : ")
					fmt.Scanln(&i)
					fmt.Print("Nomor Kursi ke : ")
					fmt.Scanln(&j)
				}
			}
		} else if jenisTiket == "Regular_Kiri" {
			fmt.Println("Pilih Tempat Duduk")
			fmt.Print("Baris Kursi ke : ")
			fmt.Scanln(&i)
			for i > 0 {
				j = 1
				if datakursiRL[i][j].noiden == 0 {
					fmt.Print("Nomor Identitas : ")
					fmt.Scanln(&noiden)
					fmt.Print("Nama Panggilan  : ")
					fmt.Scanln(&nama)
					fmt.Print("Umur            : ")
					fmt.Scanln(&umur)
					if umur >= 6 && umur < 19 {
						anak = anak + 1
						datakursiRL[i][j].noiden = noiden
						datakursiRL[i][j].nama = nama
						datakursiRL[i][j].umur = umur

						bayar = hargaTiketR * 0.6
						if i > 1 {
							bayar = bayar - ((bayar * 0.03) * float64(i))
						} else {
							bayar = bayar
						}

						datakursiRL[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi     :", kodereservasi)
						fmt.Println("Baris Tempat Duduk :", i)
						fmt.Println("Harga Tiket        :", bayar)
						total = total + bayar
						fmt.Println()

						i = 0
					} else if umur >= 19 {
						dewasa = dewasa + 1
						datakursiRL[i][j].noiden = noiden
						datakursiRL[i][j].nama = nama
						datakursiRL[i][j].umur = umur
						if i > 1 {
							bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
						} else {
							bayar = hargaTiketR
						}

						datakursiRL[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi     :", kodereservasi)
						fmt.Println("Baris Tempat Duduk :", i)
						fmt.Println("Harga Tiket        :", bayar)
						total = total + bayar
						fmt.Println()

						i = 0
					} else if umur > 0 {
						fmt.Println("Umur Belum Mencukupi")
						i = 0
					}
				} else if datakursiRL[i][j].noiden != 0 {
					for j > 0 && j < N {
						j = j + 1
						if datakursiRL[i][j].noiden == 0 {
							fmt.Print("Nomor Identitas : ")
							fmt.Scanln(&noiden)
							fmt.Print("Nama Panggilan  : ")
							fmt.Scanln(&nama)
							fmt.Print("Umur            : ")
							fmt.Scanln(&umur)
							if umur >= 6 && umur < 19 {
								anak = anak + 1
								datakursiRL[i][j].noiden = noiden
								datakursiRL[i][j].nama = nama
								datakursiRL[i][j].umur = umur

								bayar = hargaTiketR * 0.6
								if i > 1 {
									bayar = bayar - ((bayar * 0.03) * float64(i))
								} else {
									bayar = bayar
								}

								datakursiRL[i][j].kodereservasi = kodereservasi

								fmt.Println("\n   THIS IS YOUR TICKET")
								fmt.Println("Kode Reservasi     :", kodereservasi)
								fmt.Println("Baris Tempat Duduk :", i)
								fmt.Println("Harga Tiket        :", bayar)
								fmt.Println()

								i = 0
								j = 0
							} else if umur >= 19 {
								dewasa = dewasa + 1
								datakursiRL[i][j].noiden = noiden
								datakursiRL[i][j].nama = nama
								datakursiRL[i][j].umur = umur
								if i > 1 {
									bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
								} else {
									bayar = hargaTiketR
								}

								datakursiRL[i][j].kodereservasi = kodereservasi

								fmt.Println("\n   THIS IS YOUR TICKET")
								fmt.Println("Kode Reservasi     :", kodereservasi)
								fmt.Println("Baris Tempat Duduk :", i)
								fmt.Println("Harga Tiket        :", bayar)
								total = total + bayar
								fmt.Println()

								i = 0
								j = 0
							} else if umur > 0 {
								fmt.Println("Umur Belum Mencukupi")
								i = 0
								j = 0
							}
						}
						if datakursiRL[i][5].noiden != 0 {
							fmt.Print("PENUH\n\n")
							fmt.Print("Baris Kursi ke : ")
							fmt.Scanln(&i)
						}
					}
				}

			}

		} else if jenisTiket == "Regular_Kanan" {
			fmt.Println("Pilih Tempat Duduk")
			fmt.Print("Baris Kursi ke : ")
			fmt.Scanln(&i)
			for i > 0 {
				j = 1
				if datakursiRR[i][j].noiden == 0 {
					fmt.Print("Nomor Identitas : ")
					fmt.Scanln(&noiden)
					fmt.Print("Nama Panggilan  : ")
					fmt.Scanln(&nama)
					fmt.Print("Umur            : ")
					fmt.Scanln(&umur)
					if umur >= 6 && umur < 19 {
						anak = anak + 1
						datakursiRR[i][j].noiden = noiden
						datakursiRR[i][j].nama = nama
						datakursiRR[i][j].umur = umur

						bayar = hargaTiketR * 0.6
						if i > 1 {
							bayar = bayar - ((bayar * 0.03) * float64(i))
						} else {
							bayar = bayar
						}

						datakursiRR[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi     :", kodereservasi)
						fmt.Println("Baris Tempat Duduk :", i)
						fmt.Println("Harga Tiket        :", bayar)
						total = total + bayar
						fmt.Println()

						i = 0
					} else if umur >= 19 {
						dewasa = dewasa + 1
						datakursiRR[i][j].noiden = noiden
						datakursiRR[i][j].nama = nama
						datakursiRR[i][j].umur = umur
						if i > 1 {
							bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
						} else {
							bayar = hargaTiketR
						}

						datakursiRR[i][j].kodereservasi = kodereservasi

						fmt.Println("\n   THIS IS YOUR TICKET")
						fmt.Println("Kode Reservasi     :", kodereservasi)
						fmt.Println("Baris Tempat Duduk :", i)
						fmt.Println("Harga Tiket        :", bayar)
						total = total + bayar
						fmt.Println()

						i = 0
					} else if umur > 0 {
						fmt.Println("Umur Belum Mencukupi")
						i = 0
					}
				} else if datakursiRR[i][j].noiden != 0 {
					for j > 0 && j < N {
						j = j + 1
						if datakursiRR[i][j].noiden == 0 {
							fmt.Print("Nomor Identitas : ")
							fmt.Scanln(&noiden)
							fmt.Print("Nama Panggilan  : ")
							fmt.Scanln(&nama)
							fmt.Print("Umur            : ")
							fmt.Scanln(&umur)
							if umur >= 6 && umur < 19 {
								anak = anak + 1
								datakursiRR[i][j].noiden = noiden
								datakursiRR[i][j].nama = nama
								datakursiRR[i][j].umur = umur

								bayar = hargaTiketR * 0.6
								if i > 1 {
									bayar = bayar - ((bayar * 0.03) * float64(i))
								} else {
									bayar = bayar
								}

								datakursiRR[i][j].kodereservasi = kodereservasi

								fmt.Println("\n   THIS IS YOUR TICKET")
								fmt.Println("Kode Reservasi     :", kodereservasi)
								fmt.Println("Baris Tempat Duduk :", i)
								fmt.Println("Harga Tiket        :", bayar)
								total = total + bayar
								fmt.Println()

								i = 0
								j = 0
							} else if umur >= 19 {
								dewasa = dewasa + 1
								datakursiRR[i][j].noiden = noiden
								datakursiRR[i][j].nama = nama
								datakursiRR[i][j].umur = umur
								if i > 1 {
									bayar = hargaTiketR - ((hargaTiketR * 0.03) * float64(i))
								} else {
									bayar = hargaTiketR
								}

								datakursiRR[i][j].kodereservasi = kodereservasi

								fmt.Println("\n   THIS IS YOUR TICKET")
								fmt.Println("Kode Reservasi     :", kodereservasi)
								fmt.Println("Baris Tempat Duduk :", i)
								fmt.Println("Harga Tiket        :", bayar)
								total = total + bayar
								fmt.Println()

								i = 0
								j = 0
							} else if umur > 0 {
								fmt.Println("Umur Belum Mencukupi")
								i = 0
								j = 0
							}
						}
						if datakursiRR[i][5].noiden != 0 {
							fmt.Print("PENUH\n\n")
							fmt.Print("Baris Kursi ke : ")
							fmt.Scanln(&i)
						}

					}
				}

			}

		}
		p = p + 1
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
		j = 1
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
		j = 1
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
		j = 1
		i = i + 1
	}
}
func searchidentity() {}
func main() {
	menu()
	switch {
	case nomor == 1:
		reserve()
	case nomor == 2:
		searchreserve()
	case nomor == 3:
		searchidentity()
	}
}
