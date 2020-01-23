package main

import "fmt"

var arr [7]int

func main() {
	arr[0] = 7
	arr[1] = 11
	arr[2] = 8
	arr[3] = 5
	arr[4] = 2
	arr[5] = 100
	arr[6] = 9
	//sebelum berurut
	fmt.Println(arr)
	//sesudah berurut
	urut()
}

func urut() {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] >= arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
