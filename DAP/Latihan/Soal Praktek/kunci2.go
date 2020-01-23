package main 
import "fmt"
func main() {
	var n, bil1, bil2 int

	fmt.Scan(&n)

	i:=0
	for i < n {
		fmt.Scan(&bil1,&bil2)

		if (bil1+bil2)%2==0 {
			fmt.Print(bil1)
		} else {
			fmt.Print(bil2)
		}
		fmt.Println()
		i++
	}
}