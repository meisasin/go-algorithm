package main

/**
将有序数组转换为二叉搜索树

将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例1：
```
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
```
*/

/**
 确实有点看不太懂题
	error
*/
func SortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := TreeNode{Val: nums[mid]}
	if len(nums) == 1 {
		return &root
	}
	curr := &root
	for i := mid - 1; i >= 0; i-- {
		curr.Left = &TreeNode{Val: nums[i]}
		curr = curr.Left
	}
	if len(nums) == 2 {
		return &root
	}
	curr = &root
	root.Right = &TreeNode{Val: nums[len(nums)-1]}
	curr = root.Right
	for i := len(nums) - 2; i > len(nums)/2; i-- {
		curr.Left = &TreeNode{Val: nums[i]}
		curr = curr.Left
	}

	return &root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
