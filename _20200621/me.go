package main

import (
	"math"
)

/**
验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
说明：本题中，我们将空字符串定义为有效的回文串。

示例1：
```
输入: "A man, a plan, a canal: Panama"
输出: true
```

示例2：
```
输入: "race a car"
输出: false
```
*/

/**
简单题，代码写的确实有点长了，一点都不优雅，一会去找找 Go 自带的校验和转换方法，就不用自己写方法了
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSumWithMe(root *TreeNode) int {

	_, max := depth(root)
	return max
}

func depth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, math.MinInt32
	}
	l, maxL := depth(root.Left)
	r, maxR := depth(root.Right)

	curr := root.Val
	currMaxVal := max(max(curr, curr+l), curr+r)
	maxVal := max(max(currMaxVal, maxL), max(maxR, curr+r+l))

	return currMaxVal, maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
