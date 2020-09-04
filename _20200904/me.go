package main

import "strconv"

/**
257. 二叉树的所有路径

给定一个二叉树，返回所有从根节点到叶子节点的路径。
说明: 叶子节点是指没有子节点的节点。

示例：
```
输入:

   1
 /   \
2     3
 \
  5

输出: ["1->2->5", "1->3"]

解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
```
*/

/**
Error
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func BinaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	var ans []string
	lv := BinaryTreePaths(root.Left)
	rv := BinaryTreePaths(root.Right)

	if lv != nil {
		for i := 0; i < len(lv); i++ {
			v := strconv.Itoa(root.Val) + "->" + lv[i]
			ans = append(ans, v)
		}
	}
	if rv != nil {
		for i := 0; i < len(rv); i++ {
			v := strconv.Itoa(root.Val) + "->" + rv[i]
			ans = append(ans, v)
		}
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
