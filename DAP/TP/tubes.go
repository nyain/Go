package main

import "fmt"

type tiket struct {
	nomor   string
	jenis   string
	jam     int
	menit   int
	tanggal int
	bulan   int
	tahun   int
}

const t = 100
const m = 100
const n = 100

var kartu [200]tiket
var motor [t]tiket
var mobil [n][m]tiket
var isi, pilih int

var hour, min int

func bayar() {
	var hourout, minout, tarif, harga int

	fmt.Print("masukkan jam berapa anda masuk (hh mm) : ")
	fmt.Scan(&hour, &min)
	for (hour > 22 || hour < 9) || (min >= 60 || min <= 0) {
		fmt.Print("waktu masuk tidak valid, ulangi : ")
		fmt.Scan(&hour, &min)
	}
	fmt.Print("masukkan tarif per jam : ")
	fmt.Scan(&tarif)
	fmt.Print("masukkan jam berapa anda keluar (hh mm) : ")
	fmt.Scan(&hourout, &minout)
	for (hourout > 22 || hourout < 9) || (minout > 60 || minout < 0) {
		fmt.Print("waktu masuk tidak valid, ulangi : ")
		fmt.Scan(&hourout, &minout)
	}

	hourout -= hour
	minout -= min

	if minout < 11 && hourout == 0 {
		fmt.Println("parkir anda gratis")
	} else if hourout > 0 && minout <= 10 {
		harga = hourout * tarif
	} else {
		harga = (hourout + 1) * tarif
	}
	fmt.Println(harga)
}

func utama() {
	fmt.Println("1. input kendaraan. tidak penuh, masukan data kendaraan, dan kendaraan ditempatkan")
	fmt.Println("2. biaya parkir. motor keluar, rapihkan motor")
	fmt.Println("3. mencari kendaraan dengan nomor. mengeluarkan data kendaraan")
	fmt.Println("4. mencetak statistik harian : berapa kendaraan dan jenisnya pada hari itu")
	fmt.Println("5. mencetak okupansi parkir masing-masing area")
	fmt.Println("6. menghitung pendapatan sehari")
	fmt.Println("7. print mobil yang masuk dan keluar beserta nomornya")
	fmt.Println("8. exit program")
}

func main() {
	utama()
	fmt.Print("\nPilih menu: ")
	fmt.Scan(&pilih)
	switch {
	case pilih == 1:
		cekpenuh()
	case pilih == 2:
		bayar()
	case pilih == 3:
		carinomor()
	case pilih == 4:
		jumlahkendaraan()
	case pilih == 5:
		okupansi()
	case pilih == 6:
		pendapatan()
	case pilih == 7:
		printtiket()
	default:
		return
	}
}

func carinomor() {
	fmt.Println("Gua ganteng")
}

func jumlahkendaraan() {
	fmt.Println("Gua ganteng")

}

func okupansi() {
	fmt.Println("Gua ganteng")

}

func pendapatan() {
	fmt.Println("Gua ganteng")

}

func printtiket() {
	fmt.Println("Gua ganteng")
}

func cekpenuh() {
	var i int
	fmt.Print("masukan nomor, jenis, jam masuk (hh mm), tanggal, bulan, tahun = ")
	for i = 0; i < 200; i++ {
		if kartu[i].nomor == " " && kartu[i].jenis == " " && kartu[i].jam == 0 && kartu[i].menit == 0 && kartu[i].tanggal == 0 && kartu[i].bulan == 0 && kartu[i].tahun == 0 {
			fmt.Scan(&kartu[i].nomor, &kartu[i].jenis, &kartu[i].jam, &kartu[i].menit, &kartu[i].tanggal, &kartu[i].bulan, &kartu[i].tahun)
			isi++
		}
	}
	for j := 0; j < 200; j++ {
		for k := 0; k < 200; k++ {
			if kartu[i].jenis == "mobil" {
				if mobil[i][j].nomor == " " && mobil[i][j].jenis == " " && mobil[i][j].jam == 0 && mobil[i][j].menit == 0 && mobil[i][j].tanggal == 0 && mobil[i][j].bulan == 0 && mobil[i][j].tahun == 0 {
					mobil[i][k] = kartu[i]
				}
			} else if kartu[i].jenis == "bus" {
				if mobil[i][j].nomor == " " && mobil[i][j].jenis == " " && mobil[i][j].jam == 0 && mobil[i][j].menit == 0 && mobil[i][j].tanggal == 0 && mobil[i][j].bulan == 0 && mobil[i][j].tahun == 0 {
					mobil[j][k] = kartu[i]
					mobil[j+1][k] = kartu[i]
				}
			} else {
				if motor[i].nomor == " " && motor[i].jenis == " " && motor[i].jam == 0 && motor[i].menit == 0 && motor[i].tanggal == 0 && motor[i].bulan == 0 && motor[i].tahun == 0 {
					motor[k] = kartu[i]
				}
			}
		}
	}
	fmt.Print(kartu[i])
}
