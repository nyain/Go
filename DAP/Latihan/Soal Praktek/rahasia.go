package main
import ("fmt"
		"math/rand"
		)
func main () {
	var seed,kunci1, kunci2, kunci3, kunci4, kunci5 int64
	
	fmt.Scan(&seed,&kunci1,&kunci2,&kunci3,&kunci4,&kunci5)
	fmt.Println(seed)
	fmt.Println((rand.Intn(80)+74), kunci1)
	fmt.Println(kunci2,(rand.Intn(80)+74))
	fmt.Println(kunci3,(rand.Intn(80)+74))
	fmt.Println((rand.Intn(80)+74), kunci4)
	fmt.Println(kunci5,(rand.Intn(80)+74))

}