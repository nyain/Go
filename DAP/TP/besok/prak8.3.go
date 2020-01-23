package main
import "fmt"

const NMAX=1000000
var data [NMAX]int

func main() {
	var n,k int
	fmt.Scan(&n,&k)
	isiArray(n)
	if posisi(n,k)!=0 {
		fmt.Println(posisi(n,k))
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func posisi(n,k int) int {
	for i:=0 ; i<n ; i++ {
		if k==data[i] {
			return i
		}
	}
	return 0
}

func isiArray(n int) {
	var x int
	for i:=0 ; i<n ; i++ {
		fmt.Scan(&x)
		data[i]=x
	}
}