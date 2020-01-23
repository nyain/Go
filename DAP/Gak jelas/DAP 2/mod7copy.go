/*

*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

const N = 10000

var points [N]Point
var numpoints int

func jarak(p1, p2 Point) float64 {
	var jarak float64

	jarak = math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
	
	return jarak
}

func bacaTitik() {
	var x, y float64	
	fmt.Scan(&x, &y)
	numpoints = 0
	for x != 0 || y != 0 {
		points[numpoints] = Point{
			x,
			y,
		}
		numpoints++
		fmt.Scan(&x, &y)
	}
}

func ambilTitikTerdekat(p1,p2 *Point) {
	var jarakmin float64
	var i,j int
	jarakmin=99999999
	for i < numpoints {
		j=1
		for j < numpoints && j!=i {

			jarak:=jarak(points[i],points[j])
	
			if jarak<jarakmin {
				jarakmin = jarak
				*p1=points[i]
				*p2=points[j]
			}
			j++
		}
		i++
	}
}


func main() {
	var p1, p2 Point

	bacaTitik()
	ambilTitikTerdekat(&p1, &p2)
	fmt.Print("Titik terdekat adalah ", "(" , p1.x, "," , p1.y, ") dan (", p2.x, ",", p2.y , ") dengan jarak ", jarak(p1,p2))
	fmt.Println()
	fmt.Println()
	fmt.Println("NAMA : VINCENTIUS ARNOLD FRIDOLIN")
	fmt.Println("NIM : 1301190221")
}