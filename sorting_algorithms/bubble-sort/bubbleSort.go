package main

import "fmt"

func main() {
	result := bubbleSort([]float64{1, 5, 4, 6})
	fmt.Println(result)
}

func bubbleSort(arr []float64) []float64 {
	n := len(arr)
	// in remain, find the biggest one to put in the last index
	for pass:=n-1;pass>=0;pass--{
		// rise up the big one in arr
		for i:=0;i<pass;i++{
			if arr[i]>arr[i+1]{
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp	
			}
		}
	}
	return arr
}