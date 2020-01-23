package main
import "fmt"
func main() {
	var s1,s2,s3,warbiasa,i float64
	fmt.Scan(&s1,&s2,&s3)
	for s1>0 && s2>0 && s3>0 {
		if s1==100 && s2==100 && s3==100 {
			warbiasa++
			i++
			fmt.Scan(&s1,&s2,&s3)
		} else {
			fmt.Scan(&s1,&s2,&s3)
			i++
		}
	}
	fmt.Println("tim dengan predikat luar biasa ada ",warbiasa," tim dari ",i," tim")
}
