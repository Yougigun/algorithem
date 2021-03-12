// #   希尔排序是D.L.Shell于1959年提出来的一种排序算法，在这之前排序算法的时间复杂度基本
// # 都是O(n2)的，希尔排序算法是突破这个时间复杂度的第一批算法之一。插入排序的效率在某些时
// # 候是很高的，比如，记录本身就是基本有序的，只需要少量的插入操作，就可以完成整个记录集的排序
// # 工作，此时直接插入很高效。还有就是记录数比较少时，直接插入的优势也比较明显。可问题在于，
// # 两个条件本身就过于苛刻，现实中记录少或者基本有序都属于特殊情况。不过别急，有条件当然是好
// # ，条件不存在，我们创造条件，也是可以去做的。于是科学家希尔研究出了一种排序，对直接插入排
// # 序改进后可以增加效率的方法。如何让待排序的记录个数较少呢？很容易想到的就是将原本有大量记
// # 录数的记录进行分组。分割成若干个子序列，此时每个子序列待排序的记录个数就比较少了。然后在
// # 这些子序列内分别进行直接插入排序，当整个序列都基本有序时，注意只是基本有序时，再对全体记
// # 录进行一次直接插入排序。
// #  算法描述
// #   1)选定步长gap1(小于数据长度)，按步长gap1将数据分成gap1组（按列分的）
// #   2)所有距离为gap1倍数的数据放在同一组中
// #   3)在各组中进行直接插入排序
// #   4)重新取步长gap2(<gap1)重复上述的分组和排序，直到所取步长gapn=1,即所有数据
// #     放在同一组中进行直接插入排序为止。

// #   在希尔排序 开始时增量较大，分组较多，每组的记录数目少，故各组内直接插入较快，后来增量逐
// # 渐缩小，分组数逐渐减少，而各组的记录数目逐渐增多，但由于已经按 增量作为距离排过序，使文
// # 件较接近于有序状态，所以新的一趟排序过程也较快。 由于Shell排序算法 是按增量分组进行的排
// # 序，所以Shell排序算法是一种不稳定的排序算法。****
package main
import "fmt"


func main() {
	arr := []float64{1,9,5,3,4,4,4,4,10,2,3,4}
	arr = shellSort(arr)
	fmt.Println(arr)
}

func shellSort(arr []float64) []float64{
	n := len(arr)
	gap := n/2
	for gap>0 {
		// same gap, different group, ith element 
		for i := gap;i<n;i++ {
			// to put temp in corrent position with insertion sort
			temp := arr[i]
			j:=i-gap
			// insertion sort for ith element in one gap group
			for (j>0 && arr[j]>temp) {
				arr[j+gap] = arr[j]
				j = j-gap
			}
			arr[j+gap] = temp
		}
		gap = gap/2
	}
	return  arr 
}