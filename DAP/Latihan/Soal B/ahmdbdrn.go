package main
import "fmt"
func main() {
  var a1,a2,a3,b1,b2,b3 int
  fmt.Scanln(&a1,&a2,&a3)
  fmt.Scanln(&b1,&b2,&b3)
  //absolut
  ab1 := a1 > b1 && a2 > b2 && a3 > b3
  ab2 := b1 > a1 && b2 > a2 && b3 > a3
  //kompetisi
  kp1 := a1 > b1 && a2 > b2 && a3 > b3 || a1 > b1 && a2 < b2 && a3 > b3 ||
         a1 < b1 && a2 > b2 && a3 > b3 || a1 > b1 && a2 > b2 && a3 < b3 ||

         (a3 > b3 && a1 == b1 && a2 == b2) || (a1 > b1 && a2 == b2 && a3 == b3) ||
         (a2 > b2 && a1 == b1 && a3 == b3)

  kp2 := b1 > a1 && b2 > a2 && b3 > a3 || b1 > a1 && b2 < a2 && b3 > a3 ||
         b1 < a1 && b2 > a2 && b3 > a3 || b1 > a1 && b2 > a2 && b3 < a3 ||

         (b1 > a1 && b2 == a2 && b3 == a3) || (b2 > a2 && b1 == a1 && b3 == a3) ||
         (b3 > a3 && b1 == a1 && b2 == a2)

  fmt.Println("Pemenang absolut -> Ahmad?",ab1, "Badrun?",ab2)
  fmt.Println("Pemenang kompetisi -> Ahmad?",kp1, "Badrun?",kp2)

}
