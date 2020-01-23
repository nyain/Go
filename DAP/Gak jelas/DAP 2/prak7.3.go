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
			if a.r % 2 == 0 {
				a = pixel {
					r:0,
					g:0,
					b:0,
				}
			} else {
				a = pixel {
					r:max,
					g:0,
					b:0,
				}
			}
			img[i][j]=a
		}
	}
	if p == "P3" {
		p="P2"
	}
	fmt.Println(p)
	fmt.Println(w,h)
	fmt.Println(max)
	fmt.Println(img)
}
