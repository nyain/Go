package main

import "fmt"

func main() {
	const N = 10000
	var suara [N]int
	var vote [N]int
	var n, val, i, count, ind, temp int
	fmt.Scan(&n)
	for n != 0 {
		suara[i] = n
		if n > 0 && n <= 20 {
			val++
		}
		i++
		fmt.Scan(&n)
	}
	fmt.Println("Suara masuk =", i)
	fmt.Println("Suara sah =", val)

	for m := 1; m <= i; m++ {
		o := m
		for o > 0 && suara[o] < suara[o-1] {
			temp = suara[o-1]
			suara[o-1] = suara[o]
			suara[o] = temp
			o--
		}
	}

	for j := 0; j <= i; j++ { // i = banyak suara
		if suara[j] > 0 && suara[j] <= 20 { // deteksi angka valid
			conc := true
			k := 0
			for conc != false && k <= ind { // deteksi angka yang berulang
				if suara[j] == vote[k] {
					conc = false
				}
				k++
			}
			if conc == true { // dimasukkan ke array
				vote[ind] = suara[j]
				count = 1
				ind++
				for l := j + 1; l <= i; l++ { // hitung angka berulang
					if suara[l] == suara[j] {
						count++
					}
				}
				fmt.Println(suara[j], ":", count)
			}
		}
	}
}
