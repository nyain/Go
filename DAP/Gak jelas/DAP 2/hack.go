package main

import (
	"fmt"
)

func main() {
	var (
		pi, si, N float64
		i         int
	)

	fmt.Print("Masukkan suku Pertama: ")
	fmt.Scanln(&N)

	sukuN := int(N)
	pi = 0
	si = 1

	for i = 0; i < sukuN; i++ {

		if i%2 == 0 {
			pi += (1 / si)
		} else {
			pi -= (1 / si)
		}

		si = si + 2

	}

	pi *= 4
	fmt.Println("Hasil PI: ", pi)
	fmt.Println("Pada i ke: ", i)

}
