package main
import "fmt"
func main() {
	var tim, lulus int
	var s1,s2,s3,rata,batas float64

	fmt.Scan(&tim,&batas)

	for i := 0; i<tim; i++ {
		fmt.Scan(&s1,&s2,&s3)
		rata=(s1+s2+s3)/3
		if rata>batas {
			lulus++
		}
	}
	fmt.Println("Peserta dengan predikat honorable mention ada ",lulus," tim")
}
