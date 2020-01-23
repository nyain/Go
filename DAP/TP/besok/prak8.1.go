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

func hitungFrekuensi(data arr, m int, key int, frek *int) {
	for i:=0 ; i<m ; i++ {
		if data[i] == key {
			*frek++
		}
	}
	fmt.Println(*frek)
	*frek=0
}

func main() {
	var m,n,frek int

	fmt.Scan(&n)
	for i:=0 ; i<n ; i++ {
		isiArray(&data,&m)
		fmt.Print(max(data,m)," ")
		hitungFrekuensi(data,m,max(data,m),&frek)
	}
}
