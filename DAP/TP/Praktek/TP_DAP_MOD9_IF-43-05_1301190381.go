/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat konsep pemograman modular dengan gabungan fungsi,struct,dan array
		Program ini memuat metode pencarian dan pengurutan pada array
		Input berupa integer sesuai dengan batas yang ditentukan
		Output berupa hasil pengurutan dari terkecil ke terbesar dan sebaliknya
		Lalu user juga bisa mencari bilangan tertentu dengan metode pencarian binary search
		(bernilai true jika bilangan ada dan sebaliknya)
*/
package main

import "fmt"

const MAXSIZE = 10

type RecType struct {
	count, size int
}
type ArrType [MAXSIZE]RecType

var arr ArrType

func iSort(arr ArrType, nsize int) {
	for i := 0; i < nsize; i++ {
		for j := 0; j < i+1; j++ {
			if arr[i].count < arr[j].count {
				tampung := arr[j].count
				arr[j].count = arr[i].count
				arr[i].count = tampung
			}
		}
	}
	for i := 0; i < nsize; i++ {
		fmt.Print(arr[i].count, " ")
	}
}

func mSort(arr *ArrType, nsize int) {
	for i := 0; i < nsize; i++ {
		maks := i
		for j := i + 1; j < nsize; j++ {
			if arr[j].size > arr[maks].size {
				maks = j
			}
		}
		nmax := arr[maks].size
		arr[maks].size = arr[i].size
		arr[i].size = nmax
	}
}

func isFound(arr ArrType, n int, v int) bool {
	var kecil, besar, middle int
	kecil = 0
	besar = n - 1
	for besar > kecil {
		middle = (besar + kecil) / 2
		if arr[middle].count < v {
			besar--
		} else {
			kecil++
		}
	}
	if arr[middle].count == v {
		return true
	} else {
		return false
	}
}

func main() {
	var search int
	for i := 0; i < MAXSIZE; i++ {
		fmt.Scan(&arr[i].size)
	}
	for i := 0; i < MAXSIZE; i++ {
		fmt.Scan(&arr[i].count)
	}
	iSort(arr, MAXSIZE)
	mSort(&arr, MAXSIZE)
	fmt.Println()

	for i := 0; i < MAXSIZE; i++ {
		fmt.Print(arr[i].size, " ")
	}
	fmt.Println()

	fmt.Print("Apa angka yang kamu ingin cari : ")
	fmt.Scan(&search)
	fmt.Println(isFound(arr, MAXSIZE, search))
	fmt.Println("NAMA : VINCENT WILLIAMS JONATHAN")
	fmt.Println("NIM : 1301190381")
}
