package main

import "math"

/**
64. 最小路径和

给定一个包含非负整数的 `m x n` 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。

示例1：
```
输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
```
*/

func MinPathSum(grid [][]int) int {

	l := len(grid)
	ln := len(grid[0])

	if l == 0 && ln == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		for j := 0; j < ln; j++ {
			if i == 0 && j == 0 {
				continue
			}
			lv, uv := math.MaxInt32, math.MaxInt32
			if i > 0 {
				uv = grid[i-1][j] + grid[i][j]
			}
			if j > 0 {
				lv = grid[i][j-1] + grid[i][j]
			}
			grid[i][j] = min(lv, uv)
		}
	}

	return grid[l-1][ln-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
