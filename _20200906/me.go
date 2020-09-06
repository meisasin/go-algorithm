package main

/**
107. 二叉树的层次遍历 II

给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 `[3,9,20,null,null,15,7]`,
```
    3
   / \
  9  20
    /  \
   15   7
```
返回其自底向上的层次遍历为：
```
[
  [15,7],
  [9,20],
  [3]
]
```
*/

/**
... 最近怎么老忘做题啊
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrderBottom(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var queue []*TreeNode
	queue = append(queue, root)
	var ans [][]int

	for len(queue) > 0 {
		ql := len(queue)
		var ca []int
		for i := 0; i < ql; i++ {
			curr := queue[i]
			ca = append(ca, curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[ql:]
		ans = append(ans, ca)
	}

	al := len(ans)
	rans := make([][]int, al)

	for i := 0; i < al; i++ {
		rans[i] = ans[al-i-1]
	}
	return rans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
