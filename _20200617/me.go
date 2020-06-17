package main

import "math"

/**
最佳观光组合

给定正整数数组 `A`，`A[i]` 表示第 `i` 个观光景点的评分，并且两个景点 `i` 和 `j` 之间的距离为 `j - i`。
一对景点`（i < j）`组成的观光组合的得分为`（A[i] + A[j] + i - j）`：景点的评分之和减去它们两者之间的距离。
返回一对观光景点能取得的最高分。

示例：
```
输入：[8,1,5,2,6]
输出：11
解释：i = 0, j = 2, A[i] + A[j] + i - j = 8 + 5 + 0 - 2 = 11
```

提示：
- `2 <= A.length <= 50000`
- `1 <= A[i] <= 1000`
*/

/**
暴力吧，超时。感觉是用动态规划，没找到逻辑点
*/

func maxScoreSightseeingPairWithMe(A []int) int {
	max := math.MinInt64
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			score := A[i] + A[j] + i - j
			if score > max {
				max = score
			}
		}
	}
	return max
}
