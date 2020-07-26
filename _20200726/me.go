package main

/**
329. 矩阵中的最长递增路径

给定一个整数矩阵，找出最长递增路径的长度。
对于每个单元格，你可以往上，下，左，右四个方向移动。 你不能在对角线方向上移动或移动到边界外（即不允许环绕）。

示例1：
```
输入: nums =
[
  [9,9,4],
  [6,6,8],
  [2,1,1]
]
输出: 4
解释: 最长递增路径为 [1, 2, 6, 9]。
```

示例2：
```
输入: nums =
[
  [3,4,5],
  [3,2,6],
  [2,2,1]
]
输出: 4
解释: 最长递增路径是 [3, 4, 5, 6]。注意不允许在对角线方向上移动。
```
```
*/

/**
...就喜欢做这种矩阵的题
*/
var (
	rowP = []int{-1, 1, 0, 0}
	celP = []int{0, 0, 1, -1}
)

func LongestIncreasingPath(matrix [][]int) int {

	row := len(matrix)
	if row == 0 {
		return 0
	}
	cel := len(matrix[0])
	if cel == 0 {
		return 0
	}

	tmp := make([][]int, row)
	for i := 0; i < row; i++ {
		tmp[i] = make([]int, cel)
	}
	maxPathVal := 0
	for i := 0; i < row; i++ {
		for j := 0; j < cel; j++ {
			maxPathVal = max(maxPathVal, maxPath(i, j, matrix, tmp))
		}
	}
	return maxPathVal
}

func maxPath(rowAt int, celAt int, matrix [][]int, tmp [][]int) int {
	maxPathVal := 1
	for i := 0; i < 4; i++ {
		currRow := rowAt + rowP[i]
		currCel := celAt + celP[i]
		if currRow >= 0 && currCel >= 0 && currRow < len(tmp) && currCel < len(tmp[0]) && matrix[currRow][currCel] > matrix[rowAt][celAt] {
			if tmp[currRow][currCel] != 0 {
				maxPathVal = max(maxPathVal, tmp[currRow][currCel]+1)
			} else {
				maxPathVal = max(maxPathVal, maxPath(currRow, currCel, matrix, tmp)+1)
			}
		}
	}
	tmp[rowAt][celAt] = maxPathVal
	return maxPathVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
