package main

/**
寻找两个正序数组的中位数

给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
你可以假设 nums1 和 nums2 不会同时为空。

示例 1:
```
nums1 = [1, 3]
nums2 = [2]

则中位数是 2.0
```

示例 2:
```
nums1 = [1, 2]
nums2 = [3, 4]

则中位数是 (2 + 3)/2 = 2.5
```
*/

/**
执行的有点慢
*/
func findMedianSortedArraysWithMe(nums1 []int, nums2 []int) float64 {

	res := merge(nums1, nums2)
	if len(res)%2 == 0 {
		if len(res) == 2 {
			return float64(res[0]+res[1]) / 2
		}
		return float64(res[len(res)/2]+res[len(res)/2-1]) / 2
	}
	return float64(res[len(res)/2])
}

func merge(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}

	var res []int
	n1p := 0
	n2p := 0
	for {
		if nums1[n1p] < nums2[n2p] {
			res = append(res, nums1[n1p])
			n1p++
		} else {
			res = append(res, nums2[n2p])
			n2p++
		}
		if n1p == len(nums1) {
			res = append(res, nums2[n2p:len(nums2)]...)
			break
		}
		if n2p == len(nums2) {
			res = append(res, nums1[n1p:len(nums1)]...)
			break
		}
	}
	return res
}
