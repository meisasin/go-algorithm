package main

/**
二叉树的层序遍历

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。


示例：
二叉树：[3,9,20,null,null,15,7],
```
    3
   / \
  9  20
    /  \
   15   7
```
返回其层次遍历结果：
```
[
  [3],
  [9,20],
  [15,7]
]
```
*/

/**
今天的题没有中等题的味道，也是第一道直接在 leetcode 上手敲完成的，开心
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderWithMe(root *TreeNode) [][]int {

	var stack []*TreeNode
	if root != nil {
		stack = append(stack, root)
	}

	var res [][]int
	for len(stack) > 0 {
		size := len(stack)
		var currRes []int
		for i := 0; i < size; i++ {
			top := stack[i]
			currRes = append(currRes, top.Val)
			if top.Left != nil {
				stack = append(stack, top.Left)
			}
			if top.Right != nil {
				stack = append(stack, top.Right)
			}
		}
		stack = stack[size:]
		res = append(res, currRes)
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
