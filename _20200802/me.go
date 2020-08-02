package main

/**
114. 二叉树展开为链表

给定一个二叉树，原地将它展开为一个单链表。

示例1：
```
    1
   / \
  2   5
 / \   \
3   4   6
```

将其展开为：
```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```
*/

/**
为啥看见这种 TreeNode 题脑瓜子就疼
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func Flatten(root *TreeNode) {
	_, _ = sortTree(root)
}

func sortTree(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	left := root.Left
	right := root.Right
	ltail, llast := sortTree(left)
	rtail, rlast := sortTree(right)
	if root.Left != nil && root.Right != nil {
		llast.Right = rtail
		root.Left = nil
		root.Right = ltail
		return root, rlast
	}
	if root.Left != nil {
		root.Left = nil
		root.Right = ltail
		return root, llast
	}
	if root.Right != nil {
		root.Left = nil
		root.Right = rtail
		return root, rlast
	}
	return nil, nil
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
