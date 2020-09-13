package main

/**
637. 二叉树的层平均值

给定一个非空二叉树, 返回一个由每层节点平均值组成的数组。

示例1:
```
输入：
    3
   / \
  9  20
    /  \
   15   7
输出：[3, 14.5, 11]
解释：
第 0 层的平均值是 3 ,  第1层是 14.5 , 第2层是 11 。因此返回 [3, 14.5, 11] 。
```
*/

/**
今天把秋季个人赛给错过了，可惜
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func AverageOfLevels(root *TreeNode) []float64 {

	var ans []float64
	if root == nil {
		return ans
	}
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		sl := len(stack)
		sum := 0
		for i := 0; i < sl; i++ {
			sum += stack[i].Val
			if stack[i].Left != nil {
				stack = append(stack, stack[i].Left)
			}
			if stack[i].Right != nil {
				stack = append(stack, stack[i].Right)
			}
		}
		res := float64(sum) / float64(sl)
		ans = append(ans, res)

		stack = stack[sl:]
	}

	return ans
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
