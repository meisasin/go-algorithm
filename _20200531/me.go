package main

/**
对称二叉树

给定一个二叉树，检查它是否是镜像对称的。


例如，二叉树 `[1,2,2,3,4,4,3]` 是对称的。
```
    1
   / \
  2   2
 / \ / \
3  4 4  3
```

但是下面这个 `[1,2,2,null,3,null,3]` 则不是镜像对称的:
```
   1
   / \
  2   2
   \   \
   3    3
```
*/

/**
还可以用递归做，第一时间没想到
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	var stack []*TreeNode
	stack = append(stack, root.Left, root.Right)

	for len(stack) > 0 {
		l := stack[0]
		r := stack[1]
		if l == nil && r == nil {
			stack = stack[2:]
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
		stack = stack[2:]
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
