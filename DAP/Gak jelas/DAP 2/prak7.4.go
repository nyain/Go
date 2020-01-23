package main
import "fmt"

const N = 512

type pixel struct {
	r,g,b int
}
var a pixel
var img [N][N] pixel 

func main() {
	var p string
	var w,h,i,j,max int

	fmt.Scan(&p)
	fmt.Scan(&w,&h)
	fmt.Scan(&max)
	for i=0 ; w>i ; i++ {
		for j=0 ; h>j ; j++ {
			fmt.Scan(&a.r,&a.g,&a.b)
			uji:=a.r+a.g+a.b
			if uji % 2 == 1 {
				a = pixel {
					r:max,
					g:0,
					b:0,
				}
			} else {
				a = pixel {
					r:0,
					g:0,
					b:0,
				}
			}
			img[i][j]=a
		}
	}
	fmt.Println("P2")
	fmt.Println(w,h)
	fmt.Println(max)
	for i=0 ; i<w ; i++ {
		for j=0 ; j<h ; j++ {
			fmt.Print(img[i][j]," ")
		}
		fmt.Println()
	}
}