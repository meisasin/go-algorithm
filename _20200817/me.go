package main

/**
110. 平衡二叉树

给定一个二叉树，判断它是否是高度平衡的二叉树。
本题中，一棵高度平衡二叉树定义为：
一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。


示例1：
```
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
```
返回 true 。

示例2：
```
给定二叉树 [1,2,2,3,3,null,null,4,4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
```
返回 false 。
*/

/*
	简单题，一点也不简单
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	ans, _ := isBal(root, 1)
	return ans
}

func isBal(root *TreeNode, layer int) (bool, int) {
	if root.Left == nil && root.Right == nil {
		return true, layer
	}
	lm, rm := layer, layer
	lb, rb := true, true
	if root.Left != nil {
		lb, lm = isBal(root.Left, layer+1)
	}
	if root.Right != nil {
		rb, rm = isBal(root.Right, layer+1)
	}
	return lb && rb && max(lm, rm)-min(lm, rm) <= 1, max(lm, rm)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
