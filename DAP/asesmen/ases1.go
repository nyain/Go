package main
import "fmt"
func main() {
	var hdl,ldl,tri,total,rasio float32

	fmt.Print("HDL :")
	fmt.Scan(&hdl)
	fmt.Print("LDL :")
	fmt.Scan(&ldl)
	fmt.Print("triglyserida :")
	fmt.Scan(&tri)

	total=hdl+ldl+0.2*tri
	fmt.Println(total)

	rasio=total/hdl
	fmt.Println(rasio) 

	fmt.Print("sehat? ",rasio<5)
	
}