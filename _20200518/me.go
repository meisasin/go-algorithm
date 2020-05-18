package main

/**
乘积最大子数组

给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:
```
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
```

示例 2:
```
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```
*/

/**

并不简单，一遇到动态规划就懵逼

*/
func maxProductWithMe(nums []int) int {

	maxF, minF, ans := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		mx, mn := maxF, minF
		maxF = max(mx*nums[i], max(nums[i], nums[i]*mn))
		minF = min(mn*nums[i], min(nums[i], nums[i]*mx))
		ans = max(maxF, ans)
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
