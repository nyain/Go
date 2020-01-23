package main

import "fmt"
import "math/rand"

type data struct {
	nama                string
	noiden              int64
	kodereservasi, umur int
}

var datakursiV [M][N]data
var datakursiRL [K][N]data
var datakursiRR [K][N]data
var umur, i, j, kodereservasi int
var hargaTiketR, hargaTiketV, bayar float64
var anak, dewasa int
var nama, jenisTiket string
var noiden int64

const N = 6
const M = 5
const K = 6

func main() {
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiV[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi  :", kodereservasi)
					fmt.Println("No Tempat Duduk :", i, j)
					fmt.Println("Harga Tiket     :", bayar)
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiV[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi  :", kodereservasi)
					fmt.Println("No Tempat Duduk :", i, j)
					fmt.Println("Harga Tiket     :", bayar)
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiRL[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi     :", kodereservasi)
					fmt.Println("Baris Tempat Duduk :", i)
					fmt.Println("Harga Tiket        :", bayar)
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiRL[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi     :", kodereservasi)
					fmt.Println("Baris Tempat Duduk :", i)
					fmt.Println("Harga Tiket        :", bayar)
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

							rand.Seed(noiden)
							kodereservasi = rand.Intn(1000000)
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

							rand.Seed(noiden)
							kodereservasi = rand.Intn(1000000)
							datakursiRL[i][j].kodereservasi = kodereservasi

							fmt.Println("\n   THIS IS YOUR TICKET")
							fmt.Println("Kode Reservasi     :", kodereservasi)
							fmt.Println("Baris Tempat Duduk :", i)
							fmt.Println("Harga Tiket        :", bayar)
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiRR[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi     :", kodereservasi)
					fmt.Println("Baris Tempat Duduk :", i)
					fmt.Println("Harga Tiket        :", bayar)
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

					rand.Seed(noiden)
					kodereservasi = rand.Intn(1000000)
					datakursiRR[i][j].kodereservasi = kodereservasi

					fmt.Println("\n   THIS IS YOUR TICKET")
					fmt.Println("Kode Reservasi     :", kodereservasi)
					fmt.Println("Baris Tempat Duduk :", i)
					fmt.Println("Harga Tiket        :", bayar)
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

							rand.Seed(noiden)
							kodereservasi = rand.Intn(1000000)
							datakursiRR[i][j].kodereservasi = kodereservasi

							fmt.Println("\n   THIS IS YOUR TICKET")
							fmt.Println("Kode Reservasi     :", kodereservasi)
							fmt.Println("Baris Tempat Duduk :", i)
							fmt.Println("Harga Tiket        :", bayar)
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

							rand.Seed(noiden)
							kodereservasi = rand.Intn(1000000)
							datakursiRR[i][j].kodereservasi = kodereservasi

							fmt.Println("\n   THIS IS YOUR TICKET")
							fmt.Println("Kode Reservasi     :", kodereservasi)
							fmt.Println("Baris Tempat Duduk :", i)
							fmt.Println("Harga Tiket        :", bayar)
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

}
