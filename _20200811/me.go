package main

import "fmt"

/**
130. 被围绕的区域

给定一个二维的矩阵，包含 `'X'` 和 `'O'`（字母 `O`）。
找到所有被 `'X'` 围绕的区域，并将这些区域里所有的 `'O'` 用 `'X'` 填充。

示例:
```
X X X X
X O O X
X X O X
X O X X
```
运行你的函数后，矩阵变为：
```
X X X X
X X X X
X X X X
X O X X
```

解释:
被围绕的区间不会存在于边界上，换句话说，任何边界上的 `'O'` 都不会被填充为 `'X'`。
任何不在边界上，或不与边界上的 `'O'` 相连的 `'O'` 最终都会被填充为 `'X'`。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
*/

/**
整半天，还是有 BUG
ERROR
*/
func Solve(board [][]byte) {

	il, jl := len(board), len(board[0])
	iPos := []int{-1, 1, 0, 0}
	jPos := []int{0, 0, -1, 1}
	var search func(board [][]byte, i, j int) bool

	search = func(board [][]byte, i, j int) bool {
		if i < 0 || j < 0 || i >= il || j >= jl {
			return false
		}
		if board[i][j] == '-' || board[i][j] == 'X' {
			return true
		}
		board[i][j] = '-'

		ans := search(board, iPos[0]+i, jPos[0]+j) && search(board, iPos[2]+i, jPos[2]+j) && search(board, iPos[1]+i, jPos[1]+j) && search(board, iPos[3]+i, jPos[3]+j)
		if ans {
			board[i][j] = 'X'
		} else {
			board[i][j] = 'O'
		}
		return ans
	}
	for i := 0; i < il; i++ {
		for j := 0; j < jl; j++ {
			if i == 3 && j == 3 {
				fmt.Println("pause.")
			}
			search(board, i, j)
			fmt.Println(board)
		}
	}
}
