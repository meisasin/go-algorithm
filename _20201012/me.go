package main

import (
	"math"
	"sort"
)

/**
530. 二叉搜索树的最小绝对差

给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例1：
```
输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。
```

提示：
- 树中至少有 2 个节点。
- 本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
/**
	这道简单题花了我两个小时，最后还不是用递归写出来的
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GetMinimumDifference(root *TreeNode) int {

	var arr []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			arr = append(arr, root.Val)
			dfs(root.Left)
			dfs(root.Right)
		}
	}

	dfs(root)

	sort.Ints(arr)

	ans := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		ans = min(ans, arr[i]-arr[i-1])
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
