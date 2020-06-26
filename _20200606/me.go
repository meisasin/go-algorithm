package main

/**
最长连续序列

给定一个未排序的整数数组，找出最长连续序列的长度。
要求算法的时间复杂度为 O(n)。

示例1：
```
输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。
```
*/

import (
	"math"
	"sort"
)

/**
到底还是用到了 sort 函数，不过时间复杂度和空间复杂度还是很优的
*/
func LongestConsecutive(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	nums = append(nums, math.MaxInt64)
	max := 1
	begin := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			begin++
		} else if nums[i] != nums[i-1]+1 {
			max = maxN(max, i-begin)
			begin = i
		}
	}
	return max
}
func maxN(a, b int) int {
	if a > b {
		return a
	}
	return b
}
