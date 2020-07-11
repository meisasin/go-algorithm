package main

import "sort"

/**
315. 计算右侧小于当前元素的个数

给定一个整数数组 nums，按要求返回一个新数组 counts。数组 counts 有该性质： counts[i] 的值是  nums[i] 右侧小于 nums[i] 的元素的数量。

示例:
```
输入: [5,2,6,1]
输出: [2,1,1,0]
解释:
5 的右侧有 2 个更小的元素 (2 和 1).
2 的右侧仅有 1 个更小的元素 (1).
6 的右侧有 1 个更小的元素 (1).
1 的右侧有 0 个更小的元素.
```
*/

/**
困难题有点不科学啊，先暴力一波，肯定是有时间复杂度更低的解法
*/
func CountSmaller(nums []int) []int {

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		ans[i] = count
	}
	return ans
}

/**
快了一点，用桶进行统计，主要还是没看懂官方的题解
*/
func CountSmallerFast(nums []int) []int {

	countMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		countMap[nums[i]] = 1
	}
	tmp := make([]int, len(countMap))
	ans := make([]int, len(nums))
	idxMap := make(map[int]int)
	idx := 0
	for k, _ := range countMap {
		tmp[idx] = k
		idx++
	}
	sort.Ints(tmp)
	for i := 0; i < len(tmp); i++ {
		idxMap[tmp[i]] = i
	}

	countArr := make([]int, len(tmp))
	for i := len(nums) - 1; i >= 0; i-- {
		idx := idxMap[nums[i]]
		count := 0
		for j := 0; j < idx; j++ {
			count += countArr[j]
		}
		ans[i] = count
		countArr[idx] = countArr[idx] + 1
	}
	return ans
}
