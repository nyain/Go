package main
import "fmt"
func main() {
  var (
    tim,tinggi,rendah string
    a,b,skor1,skor2,skor3,rata float64
  )
  fmt.Scan(&tim,&skor1,&skor2,&skor3)
  rata = (skor1+skor2+skor3) / 3

  for tim != "AKHIR" || skor1 > 0 || skor2 > 0 || skor3 > 0 {
    if rata>a {
      a=rata
      tinggi=tim
    }else {
      b=rata
      rendah=tim
    }
    fmt.Scan(&tim,&skor1,&skor2,&skor3)
    rata = (skor1+skor2+skor3) / 3
  }
  fmt.Println("Tim dengan rerata faktor tertinggi adalah tim ",tinggi,"dengan nilai ",a)
  fmt.Println("Tim dengan rerata faktor terendah adalah tim ",rendah,"dengan nilai ",b)
}
