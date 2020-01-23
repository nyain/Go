package main
import "fmt"
func main() {
	var n,faktorial int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	faktorial=1

	for i:=0; i<n; {
		faktorial=n*faktorial
		n -= 1

	}
	fmt.Println("nilai faktorial : ",faktorial)
}
