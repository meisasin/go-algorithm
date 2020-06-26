package main

import "fmt"

/**
二叉树的最近公共祖先

给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
![-1.png](./source/-1.png)

示例 1:
```
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。
```

示例 2:
```
输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
输出: 5
解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。
```

说明:

所有节点的值都是唯一的。
p、q 为不同节点且均存在于给定的二叉树中。
*/
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

/**
好吧，今天这题有点做不出来
	我发现我对 TreeNode 有恐惧症，这个结构体是真的难用
*/
func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var pcenter []int
	var qcenter []int
	findNode(root, p, &pcenter)
	findNode(root, q, &qcenter)

	fmt.Println(pcenter)
	fmt.Println(qcenter)
	return &TreeNode{}
}

/**
查询结点
*/
func findNode(root, n *TreeNode, center *[]int) bool {

	if root == nil {
		return false
	}

	if root.Val == n.Val {
		return true
	}

	// 左边找
	if findNode(root.Left, n, center) {

		centera := *center
		centera = append(centera, root.Val)

		return true
	}
	if findNode(root.Right, n, center) {

		centera := *center
		centera = append(centera, root.Val)
		return true
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
