package main

import "fmt"

var suara, jumlah, valid int

const N = 100000

var sama [N]int
var arr [N]int
var urut[N]int

func hitungsuara() {
	fmt.Scan(&suara)
	for suara != 0 {
		if suara > 0 && suara < 21 {
			arr[jumlah] = suara
			valid++
		}
		jumlah++
		fmt.Scan(&suara)
	}
	fmt.Println("Suara masuk : ", jumlah)
	fmt.Println("Suara sah : ", valid)
}

func cekjumlah() {
	var i, j int

	for i=0 ; i<jumlah ; i++ {
		sama[i]=1
	}

	for i = 0; i < jumlah; i++ {
		for j=1 ; j<jumlah ; j++ {
			if i==j {
				j++
			}
			if arr[i]==arr[j] {
				sama[i]+=1
			}
		}
		for j=0 ; j<jumlah && arr[i] != 0 && arr[i]!=arr[j] ; j++ {
				fmt.Println(arr[i],"=",sama[i])
			}
		}

}

func main() {
	hitungsuara()
	cekjumlah()
}
