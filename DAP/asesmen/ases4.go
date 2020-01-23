/*diberikan input 3 buah bilangan interger
 positif. program fpb.go dibawah seharusanya 
 mencari faktor persekutuan terbesrar dari 
 ketiga bilangan. copy paste program berikut 
 agar dapat berjalan sesuai interaksi yang 
 diberikan
*/

package main 
import "fmt"
func main() {
	var bil1,bil2,bil3,fak int

	fmt.Print("masukan 3 bilangan : ")
	fmt.Scan(&bil1,&bil2,&bil3)

	for bil1 > 0 && bil2 > 0 && bil3>0 {
		if bil2%bil1==0 {
			fak=bil1
		} else if bil3%bil2==0 {
			fak=bil2
		} else {
			fak=bil3 
		}

		found:=false

		for !found {
			if bil1%fak==0 && bil2%fak==0 && bil3%fak==0 {
				found=true
			} else {
				fak--
			}
			
		}

		fmt.Println(fak)
		fmt.Print("masukan 3 bilangan : ")
		fmt.Scan(&bil1,&bil2,&bil3)
	}
	
}