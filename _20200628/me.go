package main

import "math"

/**
长度最小的子数组

给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的连续子数组，并返回其长度。如果不存在符合条件的连续子数组，返回 0。

示例1：
```
输入: s = 7, nums = [2,3,1,2,4,3]
输出: 2
解释: 子数组 [4,3] 是该条件下的长度最小的连续子数组。
```

进阶:
如果你已经完成了O(n) 时间复杂度的解法, 请尝试 O(n log n) 时间复杂度的解法。
*/

/**
被连续子数组给绕住了，一直以为是 [3,4,5,6,7] .... 暴力解法
*/
func MinSubArrayLen(s int, nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	min := math.MaxInt32
	for i := 0; i < l; i++ {
		sum := 0
		for j := i; j < l; j++ {
			sum += nums[j]
			// fmt.Println("i :", i, ", sum is ", sum)
			if sum >= s {
				min = minNum(min, j-i+1)
				break
			}
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

func MinSubArrayLen_On(s int, nums []int) int {

	l := len(nums)
	if l == 0 {
		return 0
	}
	begin, end, sum := 0, 0, 0
	min := math.MaxInt32
	for i := 0; i < l; i++ {
		sum += nums[i]
		end++
		if sum >= s {
			min = minNum(min, end-begin)
			for sum >= s {
				min = minNum(min, end-begin)
				sum -= nums[begin]
				begin++
			}
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

func minNum(x, y int) int {
	if x < y {
		return x
	}
	return y
}
