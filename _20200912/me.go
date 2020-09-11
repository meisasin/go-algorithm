package main

import (
	"strconv"
)

/**
216. 组合总和 III

找出所有相加之和为 n 的 k 个数的组合。组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

说明：
- `所有数字都是正整数。`
- `解集不能包含重复的组合。 `

示例1:
```
输入: k = 3, n = 7
输出: [[1,2,4]]
```

示例 2：
```
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
```
*/

/**
好吧，还算是写出来了，去重用的还是不好
*/
func CombinationSum3(k int, n int) (ans [][]int) {

	var com []int
	var dfs func(target, k, s int)

	m := map[string][]int{}
	ms := ""
	dfs = func(target, k, s int) {
		if target == 0 && k == 0 {
			if _, ok := m[ms]; !ok {
				res := append([]int{}, com...)
				ans = append(ans, res)
				m[ms] = res
			}
			return
		}
		if s > 9 || k == 0 {
			return
		}

		dfs(target, k, s+1)
		for i := s; i <= 9; i++ {
			if target-s < 0 {
				return
			}
			com = append(com, s)
			ms += strconv.Itoa(s)
			dfs(target-s, k-1, i+1)
			com = com[:len(com)-1]
			ms = ms[:len(ms)-1]
		}
	}

	dfs(n, k, 1)
	return ans
}
