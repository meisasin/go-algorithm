package main

import (
	"math"
	"sort"
)

/**
最接近的三数之和

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。
返回这三个数的和。假定每组输入只存在唯一答案。

示例1：
```
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
```

提示：
- `3 <= nums.length <= 10^3`
- `-10^3 <= nums[i] <= 10^3`
- `-10^4 <= target <= 10^4`
*/

/**
好恶心这题，边界老是整差
ERROR
*/
func ThreeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	// [-4 -2 1 2 7 16]

	min := math.MaxInt32
	minStep := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		l := i + 1
		r := len(nums) - 1
		currMin := nums[i] + nums[l] + nums[r]
		currStep := currMin - target

		// fmt.Println("~~~~~~~ i is ", i)
		for ; l < len(nums)-1; l++ {
			sum := nums[i] + nums[l+1] + nums[r]
			step := sum - target

			if step < 0 {
				step *= -1
			}
			// fmt.Println("Sum is ", sum, ", Step is ", step, ", Min is ", min, ", MinStep is ", minStep)
			if step >= currStep {
				break
			}
		}

		for ; l < r-1; r-- {
			// fmt.Println(" ===== nums[i] is ", nums[i], ", l is ", l,  ", r is ", r)

			sum := nums[i] + nums[l] + nums[r-1]
			step := sum - target
			if step < 0 {
				step *= -1
			}
			// fmt.Println("Sum is ", sum, ", Step is ", step, ", Min is ", min, ", MinStep is ", minStep)
			if step >= currStep {
				break
			}
		}

		currMin = nums[i] + nums[l] + nums[r]
		currStep = currMin - target
		if currStep < 0 {
			currStep *= -1
		}
		if currStep < minStep {
			minStep = currStep
			min = currMin
		}
		// fmt.Println("nums[i] is ", nums[i], ", l is ", l,  ", r is ", r)
		// fmt.Println("CurrMin is ", currMin, ", CurrStep is ", currStep, ", Min is ", min, ", MinStep is ", minStep)

	}
	// fmt.Println("---------------")
	return min
}
