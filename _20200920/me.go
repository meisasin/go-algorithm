package main

/**
78. 子集

给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。
说明：解集不能包含重复的子集。

示例:
```
输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
```
*/

/**
哈哈，前几天的 回溯没白做
*/
func Subsets(nums []int) [][]int {

	var ans [][]int
	ans = append(ans, []int{})
	nl := len(nums)
	if nl == 0 {
		return ans
	}

	var com []int
	var dfs func(action int)

	dfs = func(action int) {
		if action == nl {
			return
		}
		com = append(com, nums[action])
		ans = append(ans, append([]int{}, com...))
		for i := action + 1; i < nl; i++ {
			dfs(i)
		}
		com = com[:len(com)-1]
	}

	for i := 0; i < nl; i++ {
		dfs(i)
	}

	return ans
}
