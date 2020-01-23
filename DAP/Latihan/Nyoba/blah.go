package main
import "fmt"
func main() {
    var a,b,c,d,e,char1,char2,char3 rune
    fmt.Scanln(&a, &b, &c, &d, &e)
    fmt.Printf("%c%c%c%c%c\n", a,b,c,d,e)
    fmt.Scanf("%c%c%c\n",&char1,&char2,&char3)
    fmt.Printf("%c%c%c\n",char1+1,char2+1,char3+1)
}
