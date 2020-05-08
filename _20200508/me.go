package main

import (
	"fmt"
	"math"
)

/**
最大正方形

在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
示例 1:
```
输入:

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
```
*/

/**
写的效率就很慢

执行结果：
通过
显示详情
执行用时 : 68 ms , 在所有 Go 提交中击败了 5.34% 的用户
内存消耗 : 7.3 MB , 在所有 Go 提交中击败了 11.11% 的用户
*/
func maximalSquareWithMe(matrix [][]byte) int {

	// 定义一个队列，用于存储值为 1 的索引
	var queue [][]int

	for x := range matrix {
		for y, m := range matrix[x] {
			if m == 1 {
				queue = append(queue, []int{x, y})
			}
		}
	}

	maxSquare := 0
	for len(queue) > 0 {

		maxSquare = int(math.Max(float64(maxSquare), float64(recursionSquare(matrix, queue[0]))))
		queue = queue[1:]

	}

	return maxSquare

}

/**
递归计算面积
*/
func recursionSquare(matrix [][]byte, idx []int) int {

	fmt.Println(matrix)

	//fmt.Println("---------------- Index")
	fmt.Println(idx)

	sideLenght := 1

	x := idx[0]
	y := idx[1]

	xlen := len(matrix[0])
	ylen := len(matrix)

	//fmt.Println("xlen ylen ", xlen, " ", ylen)
	// 未过界
a:
	for sideLenght+x < ylen && sideLenght+y < xlen {
		//fmt.Println("------------> sideLenght ", sideLenght)
		// 循环 X 轴
		for i := 0; i < sideLenght; i++ {
			if matrix[x+sideLenght][y+i] == 0 {
				break a
			}
		}

		// 循环 Y 轴
		for i := 0; i <= sideLenght; i++ {
			if matrix[x+i][y+sideLenght] == 0 {
				break a
			}
		}

		sideLenght += 1
	}
	return sideLenght * sideLenght
}
