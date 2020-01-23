package main

import "fmt"
import "math"

var distance, distance1, print, x1, y1, x2, y2 float64

const N = 10000

var points [N]Point
var numpoints int

type Point struct {
	x, y float64
}

func jarak(p1, p2 Point) float64 {
	i := 0
	j := 1
	distance1 = 999999

	for i < numpoints {
		for j < numpoints {
			p1 = points[i]
			p2 = points[j]

			x1 = p1.x
			x2 = p2.x
			y1 = p1.y
			y2 = p2.y

			distance = math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))

			j++
			if i == j {
				j++
			}

		}
		j = 0
		i++
	}
	return print
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

func ambilTitikTerdekat(p1, p2 Point) float64 {
	if distance1 > distance {
		print = distance
		distance1 = distance
	} else {
		print = distance1
		distance = distance1
	}
	return distance
}

func main() {
	fmt.Println("NAMA : VINCENT WILLIAMS JONATHAN")
	fmt.Println("NIM : 1301190381")
	var p1, p2 Point

	bacaTitik()

	jarak(p1, p2)
	ambilTitikTerdekat(p1, p2)
	fmt.Print("Titik terdekat adalah ", "(", x1, ",", y1, ") dan (", x2, ",", y2, ") dengan jarak ", jarak(p1, p2))
}
