package main

import "fmt"

func main() {
	fmt.Println(countSort([]int{0,2,5,4}))
}

func countSort(arr []int)[]int{
	max := 0
	for _, v := range arr {
		if v > max{
			max = v
		}
	}
	carray := make([]int,max+1)
	for _, v := range arr {
		carray[v]++
	}
	pos := 0
	for i, v := range carray {
		for j := 0; j < v; j++ {
			arr[pos] = i
			pos++
		}
	}

	return arr
}