package main

import "fmt"

func main() {
	var arr [6]int
	arr[0] = 7
	arr[1] = 3
	arr[2] = 4
	arr[3] = 2
	arr[4] = 1
	arr[5] = 1
	//len jumlah array orang
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			// arr[j].umur
			if arr[j] >= arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
