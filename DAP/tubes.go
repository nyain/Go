package main

import "fmt"

const N = 5
const M = 4
const K = 5

type data struct {
	nama, penyakit       string
	id, umur, jenisPaket int
}

var paket1 [N][M]data
var paket2 [N][K]data
var paket3 [N][K]data
var umur, i, j, p, noiden, jumlah, nomor, jenisPaket int
var hargaPaket1, hargaPaket2, hargaPaket3, bayar, total float64
var nama, penyakit string
var found, terisi bool

func menu() {
	fmt.Println("MENU ")
	fmt.Println("1. Input Pasien")
	fmt.Println("2. Pemasukkan")
	fmt.Println("3. Mencari Pasien berdasarkan paket")
	fmt.Println("4. Mencari Pasien berdasarkan nama")
	fmt.Println("5. Exit")
	fmt.Print("Masukkan Nomor :")
	fmt.Scan(&nomor)
}
func inputdataidentitas() {
	fmt.Print("Nama : ")
	fmt.Scanln(&nama)
	fmt.Print("ID  : ")
	fmt.Scanln(&noiden)
	fmt.Print("Penyakit: ")
	fmt.Scanln(&penyakit)
	fmt.Print("Usia : ")
	fmt.Scanln(&umur)
}
func datapenonton() {
	if jenisPaket == 1 {
		paket1[i][j].id = noiden
		paket1[i][j].nama = nama
		paket1[i][j].umur = umur
		paket1[i][j].penyakit = penyakit
		paket1[i][j].jenisPaket = jenisPaket
		bayar = hargaPaket1
	} else if jenisPaket == 2 {
		paket2[i][j].id = noiden
		paket2[i][j].nama = nama
		paket2[i][j].umur = umur
		paket2[i][j].penyakit = penyakit
		paket2[i][j].jenisPaket = jenisPaket
		bayar = hargaPaket2
	} else if jenisPaket == 3 {
		paket3[i][j].id = noiden
		paket3[i][j].nama = nama
		paket3[i][j].umur = umur
		paket3[i][j].penyakit = penyakit
		paket3[i][j].jenisPaket = jenisPaket
		bayar = hargaPaket3
	}

	fmt.Println("\n   Struk")
	fmt.Println("Nama: ", nama)
	fmt.Println("Ruangan :", i, j)
	fmt.Println("Umur: ", umur)
	fmt.Println("Penyakit: ", penyakit)
	fmt.Println("Harga bayar     :", bayar)
	fmt.Println()
	i = 0
	j = 0
	terisi = true
	total = total + bayar
	fmt.Println()
}
func kursipenuh() {
	if jenisPaket == 1 {
		fmt.Print("PENUH\n\n")
		fmt.Print("Ruangan : ")
		fmt.Scanln(&i)
		fmt.Print("Sub-Ruangan : ")
		fmt.Scanln(&j)
	} else if jenisPaket == 2 {
		fmt.Print("PENUH\n\n")
		fmt.Print("Ruangan : ")
		fmt.Scanln(&i)
		fmt.Print("Sub-Ruangan : ")
		fmt.Scanln(&j)
	} else if jenisPaket == 3 {
		fmt.Print("PENUH\n\n")
		fmt.Print("Ruangan : ")
		fmt.Scanln(&i)
		fmt.Print("Sub-Ruangan : ")
		fmt.Scanln(&j)
	}
}

//paket2
func paketqq() {
	fmt.Print("Ruangan : ")
	fmt.Scanln(&i)
	fmt.Print("Sub-Ruangan : ")
	fmt.Scanln(&j)
	for !terisi {
		if paket2[i][j].id == 0 {
			inputdataidentitas()
			datapenonton()
		} else if paket2[i][j].id != 0 {
			kursipenuh()
		}
	}
}

//paket3
func paketrr() {
	fmt.Print("Ruangan : ")
	fmt.Scanln(&i)
	fmt.Print("Sub-Ruangan : ")
	fmt.Scanln(&j)
	for !terisi {
		if paket3[i][j].id == 0 {
			inputdataidentitas()
			datapenonton()
		} else if paket3[i][j].id != 0 {
			kursipenuh()
		}
	}
}

//input pasien
func reserve() {

	fmt.Print("Masukkan Jumlah Pasien :")
	fmt.Scanln(&jumlah)
	p = 0
	total = 0
	for p < jumlah {
		terisi = false
		bayar = 0
		hargaPaket1 = 150000
		hargaPaket2 = 100000
		hargaPaket3 = 90000
		fmt.Print("Jenis Paket(Paket 1/Paket 2/Paket 3) :")
		fmt.Scanln(&jenisPaket)
		if jenisPaket == 1 || jenisPaket == 2 || jenisPaket == 3 {
			if jenisPaket == 1 {
				fmt.Print("Ruangan : ")
				fmt.Scanln(&i)
				fmt.Print("Sub-Ruangan : ")
				fmt.Scanln(&j)
				for !terisi {
					if paket1[i][j].id == 0 {
						inputdataidentitas()
						datapenonton()
					} else if paket1[i][j].id != 0 {
						kursipenuh()
					}
				}
			} else if jenisPaket == 2 {
				paketqq()

			} else if jenisPaket == 3 {
				paketrr()
			}
			p++
		} else {
			fmt.Println("Data salah")
		}
	}
	menu()
}
func searchidentity() {

	i = 0
	j = 0
	fmt.Print("\nMasukkan Nama Pasien :")
	fmt.Scanln(&nama)
	for i < N {
		for j < M {
			if paket1[i][j].nama == nama {
				fmt.Println("Nomor Identitas : ", paket1[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket1[i][j].nama)
				fmt.Println("Umur            : ", paket1[i][j].umur)
				fmt.Println("Penyakit  : ", paket1[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 1")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if paket2[i][j].nama == nama {
				fmt.Println("Nomor Identitas : ", paket2[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket2[i][j].nama)
				fmt.Println("Umur            : ", paket2[i][j].umur)
				fmt.Println("Penyakit  : ", paket2[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 2")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if paket3[i][j].nama == nama {
				fmt.Println("Nomor Identitas : ", paket3[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket3[i][j].nama)
				fmt.Println("Umur            : ", paket3[i][j].umur)
				fmt.Println("Penyakit  : ", paket3[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 3")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
	menu()
}
func paket() {

	i = 0
	j = 0
	fmt.Print("\nMasukkan Jenis Paket :")
	fmt.Scanln(&jenisPaket)
	for i < N {
		for j < M {
			if paket1[i][j].jenisPaket == jenisPaket {
				fmt.Println("Nomor Identitas : ", paket1[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket1[i][j].nama)
				fmt.Println("Umur            : ", paket1[i][j].umur)
				fmt.Println("Penyakit  : ", paket1[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 1")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if paket2[i][j].jenisPaket == jenisPaket {
				fmt.Println("Nomor Identitas : ", paket2[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket2[i][j].nama)
				fmt.Println("Umur            : ", paket2[i][j].umur)
				fmt.Println("Penyakit  : ", paket2[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 2")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		for j < K {
			if paket3[i][j].jenisPaket == jenisPaket {
				fmt.Println("Nomor Identitas : ", paket3[i][j].id)
				fmt.Println("Nama Panggilan  : ", paket3[i][j].nama)
				fmt.Println("Umur            : ", paket3[i][j].umur)
				fmt.Println("Penyakit  : ", paket3[i][j].penyakit)
				fmt.Println("Jenis Paket     : Paket 3")
				fmt.Println("Ruangan     : ", i, j)
				fmt.Println()
			}
			j = j + 1
		}
		j = 0
		i = i + 1
	}
	menu()
}

func main() {
	menu()
	for nomor >= 1 && nomor <= 4 {
		switch {
		case nomor == 1:
			reserve()
		case nomor == 2:
			fmt.Println("euy")
		case nomor == 3:
			paket()
		case nomor == 4:
			searchidentity()
		}
	}
}
