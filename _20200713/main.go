package main

import (
	"sort"
)

func main() {

}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}
