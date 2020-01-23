/*  NIM: 1301190381
    Nama: Vincent Williams Jonathan

    Program ini memuat konsep pemograman modular dengan gabungan fungsi,struct,dan array
		Program menerima input dari user berupa string, integer, dan float64 dan input tersebut berulang sampai 5 kali
		Output program menampilkan nilai float64 terbesar, index nilai integer terkecil
		Program juga memiliki fungsi untuk mencari string dari data array, mengoutput true jika string tersebut ada
		Lalu, program juga memiliki fungsi untuk mengurutkan data array
		Jika data array terurut membesar, maka program juga akan mengoutput index array dari string yang kita cari
		Jika tidak terurut membesar dan string tidak ada di dalam data array, maka program akan mengoutput -1

*/
package main

import "fmt"

type RecType struct {
	f1 string
	f2 int
	f3 float64
}

const N = 5

type ArrType [N]RecType

var data ArrType

func rmax(data ArrType) float64 {
	var realmaks float64
	i := 0
	for N > i {
		if realmaks < data[i].f3 {
			realmaks = data[i].f3
		}
		i++
	}
	return realmaks
}

func imin(data ArrType) int {
	var intmin, index int
	intmin = 2019
	i := 0
	for N > i {
		if intmin > data[i].f2 {
			index = i
			intmin = data[i].f2
		}
		i++
	}
	return index
}

func found(data ArrType, key string) bool {
	i := 0
	for i < N {
		if key == data[i].f1 {
			return true
		}
		i++
	}
	return false
}

func pos(data ArrType, key string) int {
	var middle int
	little := 0
	big := len(data) - 1

	for little <= big {
		middle = (little + big) / 2

		if data[middle].f1 < key {
			little = middle + 1
		} else {
			big = middle - 1
		}
	}
	if key == data[middle].f1 {
		return middle
	} else {
		return -1
	}
}

func main() {
	fmt.Println("NAMA : VINCENT WILLIAMS JONATHAN")
	fmt.Println("NIM : 1301190381")
	var name string
	i := 0
	for i < N {
		fmt.Print(i+1, ".) input string, integer, float64 berulang sebanyak ", N, " kali : ")
		fmt.Scan(&data[i].f1, &data[i].f2, &data[i].f3)
		i++
	}
	fmt.Println("nilai float terbesar dalam array : ", rmax(data))
	fmt.Println("index di mana integer terkecil dalam array : ", imin(data))
	fmt.Print("masukan string yang ingin dicari : ")
	fmt.Scan(&name)
	fmt.Println("apakah string tersebut ada dalam array ? ", found(data, name))
	fmt.Println("pada index ke berapa string tersebut berada .. : ", pos(data, name))
}
