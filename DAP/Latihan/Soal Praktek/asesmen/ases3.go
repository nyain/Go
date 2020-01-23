package main
import "fmt"
func main() {
	var a,b,c int
	fmt.Print("Sisi segitiga : ")
	fmt.Scan(&a,&b,&c)

	if (a<c || b<c) && (a>0 && b>0 && c>0){
		if a*a+b*b>c*c {
			if a==b || a==c {
					fmt.Println("Segitiga sama kaki")
			} else {
				fmt.Println("Segitiga sembarang")
			}

		} else if a*a+b*b==c*c {
			fmt.Println("Segitiga siku-siku")
		} else {
			fmt.Println("Bukan segitiga")
		}
	} else if a==b && a==c {
		fmt.Println("Segitiga sama sisi")
	} else {
		fmt.Println("Data salah")
	}
}
