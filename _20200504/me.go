package main

import (
	"math"
)

/**
跳跃游戏 II

给定一个非负整数数组，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。

示例 1:
```
输入: [2, 3, 1, 1, 4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
```

说明:
假设你总是可以到达数组的最后一个位置。
*/

/**
循环到 len(nums) - 1, 最后一个元素不访问，因为到达最后一个元素之后，就不会再往后跳跃了
*/
func Jump(nums []int) int {

	maxPosition := 0
	end := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {

		maxPosition = int(math.Max(float64(maxPosition), float64(i+nums[i])))
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

/**
没做出来，继续努力
*/
// 从最后一个元素往前跳，走到最开始的位置
func toJump(nums []int, nextStep int, idx int, countJump int) int {
	// 如果跳到最开始，且开始的步数刚好是 step ，返回
	if idx < 0 {
		return math.MaxInt32
	}
	if idx == 0 {
		if nums[idx] == nextStep {
			countJump++
			return countJump
		} else {
			return math.MaxInt32
		}
	}

	sel := math.MaxInt32
	// 判断选还是不选，如果 nums[idx] == nextStep 可选
	if nums[idx] == nextStep {
		countJump++
		idx--
		sel = toJump(nums, 1, idx, countJump) // 选
	}

	nextStep++
	idx--
	nosel := toJump(nums, nextStep, idx, countJump) // 不选

	return int(math.Min(float64(sel), float64(nosel)))
}
