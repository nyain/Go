package main
import "fmt"

const N=2000
type arr [N]int
var data arr

func isiArray(data *arr, m *int) {
	var x int
	fmt.Scan(&x)
	for *m=0 ; x!=999 ; *m++ {
		data[*m]=x
		fmt.Scan(&x)
	}
}

func max(data arr, m int) int {
	var max int
	for i:=0 ; i<m ; i++ {
		if max<data[i] {
			max = data[i]
		}
	}
	return max
}

func min(data arr, m int) int {
	var min int = 99999999
	for i:=0 ; i<m ; i++ {
		if min>data[i] {
			min = data[i]
		}
	}
	return min
}

func hitungFrekuensi(data arr, m int, keymax int, keymin int, z *int) {
	for i:=0 ; i<m ; i++ {
		if data[i]<keymax && data[i]>keymin {
			*z++
		}
	}
	fmt.Println(*z)
	*z=0
}

func main() {
	var m,n,z int

	fmt.Scan(&n)
	for i:=0 ; i<n ; i++ {
		isiArray(&data,&m)
		fmt.Print(min(data,m)," ",max(data,m)," ")
		hitungFrekuensi(data,m,max(data,m),min(data,m),&z)
	}
}
