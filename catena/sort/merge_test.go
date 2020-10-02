package sort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	// [1,2,3,0,0,0]
	//3
	//[2,5,6]
	//3
	ans := []int{1, 2, 3, 0, 0, 0}
	Merge(ans, 3, []int{2, 5, 6}, 3)
	fmt.Print(ans)
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	fmt.Println(nums1)
	n1l, n2l := len(nums1), len(nums2)
	i, j := n1l-1, m-1
	for j >= 0 {
		nums1[i] = nums1[j]
		i--
		j--
	}
	n1b, n2b := n1l-m, 0
	point := 0

	for point < m+n {
		if n1b == n1l {
			for ; n2b < n2l; n2b++ {
				nums1[point] = nums2[n2b]
				point++
			}
			break
		}
		if n2b == n2l {
			for ; n1b < n1l; n1b++ {
				nums1[point] = nums2[n1b]
				point++
			}
			break
		}
		if nums1[n1b] < nums2[n2b] {
			nums1[point] = nums1[n1b]
			n1b++
		} else {
			nums1[point] = nums2[n2b]
			n2b++
		}
		point++
	}
}
