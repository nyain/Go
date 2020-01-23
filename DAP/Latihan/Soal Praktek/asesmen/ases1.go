package main
import "fmt"
func main() {
	var hdl,ldl,tri float64

	fmt.Print("HDL: ")
	fmt.Scan(&hdl)
	fmt.Print("LDL: ")
	fmt.Scan(&ldl)
	fmt.Print("Triglyserida: ")
	fmt.Scan(&tri)

	fmt.Println("Total:",hdl+ldl+0.2*tri)
	fmt.Println("Rasio total/HDL:",(hdl+ldl+0.2*tri)/hdl)
	fmt.Println("sehat?",((hdl+ldl+0.2*tri)/hdl)<5)

}
