package main

/**
94. 二叉树的中序遍历

给定一个二叉树，返回它的中序 遍历。

示例:
```
输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,3,2]
```

进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/

/**
是我对 slice 有什么误解吗，为什么得传个指针进去
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func InorderTraversal(root *TreeNode) []int {

	var ans []int

	traversal(root, &ans)

	return ans
}

func traversal(root *TreeNode, bucket *[]int) {
	if root == nil {
		return
	}

	traversal(root.Left, bucket)

	*bucket = append(*bucket, root.Val)

	traversal(root.Right, bucket)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
