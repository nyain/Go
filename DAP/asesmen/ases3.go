package main
import "fmt"
func main() {
	var a,b,c int 
	fmt.Print("Sisi segitiga : ")
	fmt.Scan(&a,&b,&c)

	if (a<c || b<c) && (a>0 && b>0 && c>0) {
		if a*a+b*b>c*c {
	
			if a==b || b==c || a==c {
					fmt.Print("ini sama kaki")
			} else {
				fmt.Print("ini segitiga sembarang")
			}

		} else if a*a+b*b==c*c {
			fmt.Print("ini segitiga siku-siku")
		} else {
			fmt.Print("bukan segitiga")
		}
	} else if a==b && b==c && a==c {
		fmt.Print("sama sisi")

	} else {
		fmt.Print("bego lu")
	}
}