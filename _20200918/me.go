package main

import "sort"

/**
47. 全排列 II

给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:
```
输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
```
*/

/**
看着就像回溯算法，就是没写出来
*/
func PermuteUnique(nums []int) (ans [][]int) {

	sort.Ints(nums)
	nl := len(nums)
	var temp []int
	visit := make([]bool, nl)

	var dfs func(idx int)
	dfs = func(idx int) {
		if idx == nl {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := 0; i < nl; i++ {
			if visit[i] || (i > 0 && !visit[i-1] && nums[i] == nums[i-1]) {
				continue
			}
			temp = append(temp, nums[i])
			visit[i] = true
			dfs(idx + 1)
			visit[i] = false
			temp = temp[:len(temp)-1]
		}
	}

	dfs(0)

	return ans
}
