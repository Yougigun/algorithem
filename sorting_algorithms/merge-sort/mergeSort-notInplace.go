// 	Merge sort is a divide-and-conquer algorithm based on the idea of breaking 
// down a list into several sub-lists until each sublist consists of a single 
// element and merging those sublists in a manner that results into a sorted list.

package main

import "fmt"


func main() {
	arr := []float64{5, 4, 1, 6, 8, 3, 4}
	arr = mergeSort(arr)
	fmt.Println(arr)
}

func merge(arr1 []float64, arr2 []float64) []float64 {
	n := len(arr1)
	m := len(arr2)
	result := []float64{}
	x := 0
	y := 0
	for i := 0; i < n+m; i++ {
		if x==n  {
			result = append(result,arr2[y:]...)
			return result
		}else if y==m{
			result = append(result,arr1[x:]...)
			return result 
		}
		if arr1[x] < arr2[y] {
			result = append(result, arr1[x])
			x++
		} else {
			result = append(result, arr2[y])
			y++
		}
	}
	return result
}

func mergeSort(arr []float64) []float64 {
	if len(arr)<=1{
		return arr
	}
	arr = merge(
			mergeSort(arr[:(len(arr)/2)]),
			mergeSort(arr[(len(arr)/2):]))
	return arr
}
