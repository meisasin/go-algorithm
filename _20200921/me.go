package main

/**
538. 把二叉搜索树转换为累加树

给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，
使得每个节点的值是原来的节点值加上所有大于它的节点值之和。


例如：
```
输入: 原始二叉搜索树:
              5
            /   \
           2     13

输出: 转换为累加树:
             18
            /   \
          20     13
```
*/

/**
二叉树确实挺烦的好吧
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func ConvertBST(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, val int) int {
	ans := root.Val
	root.Val = root.Val + val
	if root.Right == nil && root.Left == nil {
		return ans
	}

	if root.Right != nil {
		rv := dfs(root.Right, val)
		ans += rv
		root.Val += rv
	}
	if root.Left != nil {
		lv := dfs(root.Left, root.Val)
		ans += lv
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
