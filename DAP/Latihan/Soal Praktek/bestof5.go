package main 
import "fmt"
func main() {
	var n1, n2 string
	var kategori, nilai1, nilai2, nilai3, soal1, soal2 int

	fmt.Scan(&kategori,&n1,&n2)
	i:=0
	for i<kategori {
		fmt.Scan(&nilai1,&nilai2,&nilai3)
		if nilai1!=0 || nilai2!=0 || nilai3!=0 {
			soal1++	
		}
		i++
	}
	
	i=0
	for i<kategori {
		fmt.Scan(&nilai1,&nilai2,&nilai3)
		if nilai1!=0 || nilai2!=0 || nilai3!=0 {
			soal2++
		}
		i++
	}
	if soal1>soal2 {
		fmt.Println(n1,"pemenang dengan menyelesaikan",soal1,"dari", kategori, "soal.")
	} else {
		fmt.Println(n2,"pemenang dengan menyelesaikan",soal2,"dari", kategori, "soal.")
	}

}