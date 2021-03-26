package quickSort

import (
	"testing"
)


func TestQuickSort(t *testing.T) {
	arr := []int{2,1,3,4,5,6}
	ans := []int{1,2,3,4,5,6}
	size := len(ans)
	res := quickSort(arr)
	sizeRes := len(res)

	if size != sizeRes {
		t.Error("For","Size","Expected",size,"Got",sizeRes)
	}

	for i, v := range ans {
		if v!=res[i]{
			t.Error("For","Value","Expected",v,"Got",res[i])
		}
	}

}