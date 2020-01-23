package main
import "fmt"
func main() {
	var nilai1,nilai2,nilai3,soal1,soal2 int
	var nama1, nama2 string
	
	i:=0
	fmt.Scan(&nama1,&nama2)
	for i<5 {
		fmt.Scan(&nilai1,&nilai2,&nilai3)
		if nilai1!=0 || nilai2!=0 || nilai3!=0 {
			soal1++	
		}
		fmt.Scan(&nilai1,&nilai2,&nilai3)
		if nilai1!=0 || nilai2!=0 || nilai3!=0 {
			soal2++	
		}
		i++
	}
	if soal1>soal2 {
		fmt.Println(nama1,"pemenang dengan menyelesaikan",soal1,"dari 5 soal.")
	} else {
		fmt.Println(nama2,"pemenang dengan menyelesaikan",soal2,"dari 5 soal.")
	}
}