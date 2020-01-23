package main

import "fmt"

func main() {
	var n, bil1, bil2 int

	fmt.Scan(&n)
	fmt.Scan(&bil1,&bil2)

	for bil1>-1 || bil2>-1 {
		
		if bil1%n ==0 {
			fmt.Print(bil2)

		} else if bil2%n ==0 {
			fmt.Print(bil1)
		}
		fmt.Println()
		fmt.Scan(&bil1,&bil2)
	}
}
