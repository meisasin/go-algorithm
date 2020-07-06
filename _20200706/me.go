package main

/**
不同路径 II

一个机器人位于一个 m x n 网格的左上角 （起始点在下图中标记为“Start” ）。
机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角（在下图中标记为“Finish”）。
现在考虑网格中有障碍物。那么从左上角到右下角将会有多少条不同的路径？

![...](./other/1.png)

网格中的障碍物和空位置分别用 1 和 0 来表示。
*说明*：m 和 n 的值均不超过 100。

示例1：
```
输入:
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
输出: 2
解释:
3x3 网格的正中间有一个障碍物。
从左上角到右下角一共有 2 条不同的路径：
1. 向右 -> 向右 -> 向下 -> 向下
2. 向下 -> 向下 -> 向右 -> 向右
```
*/

/**
处理边界确实还挺烦的，还好顺利做出来了，主要还是思路
*/
func UniquePathsWithObstacles(o [][]int) int {

	l := len(o)
	if l == 0 {
		return 0
	}
	lr := len(o[0])
	if lr == 0 {
		return 0
	}
	if l == 1 && lr == 1 {
		return o[0][0] ^ 1
	}
	if o[0][0] == 1 {
		return 0
	}
	o[0][0] = 1
	for i := 1; i < lr; i++ {
		if o[0][i] == 1 {
			o[0][i] = 0
		} else {
			o[0][i] = o[0][i-1]
		}
	}
	for i := 1; i < l; i++ {
		if o[i][0] == 1 {
			o[i][0] = 0
		} else {
			o[i][0] = o[i-1][0]
		}
	}
	for i := 1; i < l; i++ {
		for j := 1; j < lr; j++ {
			if o[i][j] == 1 {
				o[i][j] = 0
			} else {
				top, left := 0, 0
				top = o[i-1][j]
				left = o[i][j-1]

				o[i][j] = top + left
			}
		}
	}
	return o[l-1][lr-1]
}
