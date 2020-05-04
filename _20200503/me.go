package main

import "math"

/**
最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例 1:
```
输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

进阶:
如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。
*/

/**
今天的简单题有点不简单
*/
func maxSubArrayWithMe(nums []int) int {

	// 之前总和
	var sum int
	// 之前最大和
	maxSum := nums[0]

	for _, n := range nums {
		if sum+n <= 0 {
			maxSum = int(math.Max(float64(maxSum), float64(n)))
			sum = 0
		} else {
			sum += n
			maxSum = int(math.Max(float64(maxSum), float64(sum)))
		}
	}
	return maxSum
}

func maxSubArrayWithMe_(nums []int) int {

	maxSum := nums[0]
	// [-2, 1, -3, 4, -1, 2, 1, -5, 4]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i] + nums[i-1]
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}
	return maxSum
}
