package main

import "fmt"

func main() {
	A := []int{5, 3, 2, 1, 100}
	A = insertionSort(A)
	fmt.Println(A)
}

func insertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		cValue := arr[i]
		position := i
		for (position>0) && (arr[position-1] > cValue){
			arr[position] = arr[position-1]
			position = position -1 
		}
		arr[position]=cValue
	}

	return arr
}