package quickSort


func quickSort(arr []int) []int {
	// Careful
	if len(arr) <= 1  {
		return arr
	}
	arr1 := []int{}
	arr2 := []int{}
	minIndex := len(arr)/2
	mid := arr[minIndex]
	// Careful 
	arr = append(arr[:minIndex],arr[minIndex:]...)
	for _, v := range arr {
		if v < mid {
			arr1 = append(arr1, v)
		} else if v > mid {
			arr2 = append(arr2, v)
		}
	}

	return append(append(quickSort(arr1),mid),quickSort(arr2)...)
}