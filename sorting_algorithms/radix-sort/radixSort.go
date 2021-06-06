package main

import (
	"fmt"
)

func main() {
	fmt.Println(radixSort([]int{0, 2, 5, 4}))
}

func radixSort(arr []int) []int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
    // # Do counting sort for every digit. Note that instead
    // # of passing digit number, exp is passed. exp is 10^i
    // # where i is current digit number
    exp := 1
    for max / exp > 0{
        countSort(arr, exp)
        exp = exp * 10
	}
	return arr

}

func countSort(arr []int, exp int) {
	n := len(arr)
	// The output array elements that will have sorted arr
	output := make([]int, n)

	// Initialize count array as 0
	count := make([]int, 10)

	// Store count of occurrences in count[]
	for i := 0; i < n; i++ {
		index := (arr[i] / exp)
		count[index%10]++
	}

	// # Change count[i] so that count[i] now contains actual
    // # position of this digit in output array
    for i:=1; i < 10 ; i++{
        count[i] += count[i - 1]
	}

	// Build the output array
	i := n - 1
	for i >= 0 {
		index := (arr[i] / exp)
		output[count[int(index%10)]-1] = arr[i]
		count[int(index%10)] -= 1
		i -= 1
	}

	// # Copying the output array to arr[],
	// # so that arr now contains sorted numbers
	i = 0
	for i:=range(arr){
	    arr[i] = output[i]
	}
}
