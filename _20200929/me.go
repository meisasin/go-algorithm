package main

/**
145. 二叉树的后序遍历

给定一个二叉树，返回它的 后序 遍历。

示例:
```
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [3,2,1]
```
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
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
func PostorderTraversal(root *TreeNode) (ans []int) {

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			dfs(root.Right)
			ans = append(ans, root.Val)
		}
	}

	dfs(root)
	return
}

func PostorderTraversal_(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	stack = append(stack, root)
	var ans []int
	// 访问过的节点
	visit := map[*TreeNode]bool{}

	for len(stack) > 0 {
		last := stack[len(stack)-1]
		if last.Left != nil {
			if _, ok := visit[last.Left]; !ok {
				stack = append(stack, last.Left)
				visit[last.Left] = true
				continue
			}
		}
		if last.Right != nil {
			if _, ok := visit[last.Right]; !ok {
				stack = append(stack, last.Right)
				visit[last.Right] = true
				continue
			}
		}
		ans = append(ans, last.Val)
		stack = stack[:len(stack)-1]
	}
	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
