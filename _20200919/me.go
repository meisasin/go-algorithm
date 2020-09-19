package main

/**
404. 左叶子之和

计算给定二叉树的所有左叶子之和。

示例:
```
    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
```
*/

/**
...
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func SumOfLeftLeaves(root *TreeNode) int {

	var dfs func(root *TreeNode, potion int) int
	dfs = func(root *TreeNode, potion int) int {
		if root == nil {
			return 0
		}
		if root.Right == nil && root.Left == nil && potion == 0 {
			return root.Val
		}
		return dfs(root.Left, 0) + dfs(root.Right, 1)
	}

	return dfs(root, 1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
