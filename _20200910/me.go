package main

import (
	"sort"
	"strconv"
)

/**
40. 组合总和 II

给定一个数组 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。
`candidates` 中的每个数字在每个组合中只能使用一次。

说明：
- `所有数字（包括目标数）都是正整数。`
- `解集不能包含重复的组合。 `

示例1:
```
输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
```

示例 2：
```
输入: candidates = [2,5,2,1,2], target = 5,
所求解集为:
[
  [1,2,2],
  [5]
]
```
*/

/**
昨天刚做了回溯算法，今天又忘了，😓 >>> 不管怎么样，也都是自己写出来的
*/
func CombinationSum2(candidates []int, target int) [][]int {

	sort.Ints(candidates)
	var ans [][]int
	cl := len(candidates)

	m := map[string][]int{}

	var com []int
	coms := ""

	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == cl {
			return
		}

		dfs(target, idx+1)

		if target >= candidates[idx] {
			com = append(com, candidates[idx])
			last := strconv.Itoa(candidates[idx])
			coms += last

			if target == candidates[idx] {
				if _, ok := m[coms]; !ok {
					m[coms] = append([]int{}, com...)
					ans = append(ans, m[coms])
				}
			} else {
				dfs(target-candidates[idx], idx+1)
			}

			com = com[:len(com)-1]
			coms = coms[:len(coms)-len(last)]
		}
	}

	dfs(target, 0)
	return ans
}
