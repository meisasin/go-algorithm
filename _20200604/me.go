package main

/**
除自身以外数组的乘积

给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例1：
```
输入: [1,2,3,4]
输出: [24,12,8,6]
```

提示：题目数据保证数组之中任意元素的全部前缀元素和后缀（甚至是整个数组）的乘积都在 32 位整数范围内。

说明: 请不要使用除法，且在 O(n) 时间复杂度内完成此题。

进阶：
你可以在常数空间复杂度内完成这个题目吗？（ 出于对空间复杂度分析的目的，输出数组不被视为额外空间。）
*/

/**
  不能用除法 ~~~
*/
func ProductExceptSelf(nums []int) []int {

	var zeroIdx []int
	var arr = make([]int, len(nums))
	max := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroIdx = append(zeroIdx, i)
		} else {
			max *= nums[i]
		}
	}

	if len(zeroIdx) > 1 {
		return arr
	}
	if len(zeroIdx) == 1 {
		for i := 0; i < len(arr); i++ {
			if nums[i] == 0 {
				arr[i] = max
			} else {
				arr[i] = 0
			}
		}
		return arr
	}
	for i := 0; i < len(arr); i++ {
		arr[i] = max / nums[i]
	}
	return arr
}
