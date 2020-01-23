package main
import "fmt"
func main() {
	var n,faktorial int

	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	faktorial=1
	i:=0
	for i<n {
		faktorial=n*faktorial
		n--
		
	}
	fmt.Print("nilai faktorial : ",faktorial)
}