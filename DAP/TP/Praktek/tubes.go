package main

import "fmt"

var vip = [3][4]float64{}
var regkiri = [5][4]float64{}
var regkanan = [5][4]float64{}

func mewah() {
	for i := 0; i < len(vip); i++ {
		fmt.Println(vip[i])
	}
}
func kiri() {
	for j := 0; j < len(regkiri); j++ {
		fmt.Println(regkiri[j])
	}
}
func kanan() {
	for k := 0; k < len(regkanan); k++ {
		fmt.Println(regkanan[k])
	}
}

func main() {
	var kategori string
	fmt.Print("Pilih kategori: ")
	fmt.Scan(&kategori)
	if kategori == "VIP" {
		mewah()
		fmt.Println()
	} else if kategori == "regkanan" {
		kanan()
	} else {
		kiri()
	}
}
