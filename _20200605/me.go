package main

import "math"

/**
顺时针打印矩阵

输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

示例1：
```
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
```

示例1：
```
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]
```


限制：

- `0 <= matrix.length <= 100`
- `0 <= matrix[i].length <= 100`
*/

/**
  发现了一点，简单题从来不套路，中等题从来都是套路，困难题肯定要用到高级算法
*/
func spiralOrderWithMe(matrix [][]int) []int {

	point := []int{0, 0}

	var ans []int
	horzi := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	horziP := 0
	if len(matrix) == 0 {
		return ans
	}
	for point[0] >= 0 && point[1] >= 0 && point[0] < len(matrix) && point[1] < len(matrix[0]) {
		ans = append(ans, matrix[point[0]][point[1]])

		matrix[point[0]][point[1]] = math.MaxInt64
		y := point[0] + horzi[horziP][0]
		x := point[1] + horzi[horziP][1]
		if y >= len(matrix) || y < 0 || x < 0 || x >= len(matrix[0]) || matrix[y][x] == math.MaxInt64 {

			horziP++
			if horziP >= len(horzi) {
				horziP = 0
			}
			y = point[0] + horzi[horziP][0]
			x = point[1] + horzi[horziP][1]
			if y >= len(matrix) || y < 0 || x < 0 || x >= len(matrix[0]) || matrix[y][x] == math.MaxInt64 {
				break
			}
		}
		point[0] = y
		point[1] = x
	}
	return ans
}
