package main

/**
另一个树的子树

给定两个非空二叉树 s 和 t，检验 s 中是否包含和 t 具有相同结构和节点值的子树。s 的一个子树包括 s 的一个节点和这个节点的所有子孙。s 也可以看做它自身的一棵子树。

示例 1:
```
给定的树 s:
     3
    / \
   4   5
  / \
 1   2
给定的树 t：
   4
  / \
 1   2
```
返回 true，因为 t 与 s 的一个子树拥有相同的结构和节点值。

示例 1:
```
给定的树 s：
     3
    / \
   4   5
  / \
 1   2
    /
   0
给定的树 t：
   4
  / \
 1   2
```
返回 false。

```
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
看来确实是道简单的题，自己写的居然和官网第一种算法大差不差，其实第二种使用 KMP 实现的也想到了，闲太麻烦了没用
*/
func IsSubtree(s *TreeNode, t *TreeNode) bool {

	return searchAndCheck(s, t)

}

func searchAndCheck(s *TreeNode, t *TreeNode) bool {

	if s == nil {
		return false
	}

	if s.Val == t.Val && check(s, t) {
		return true
	}

	return searchAndCheck(s.Left, t) || searchAndCheck(s.Right, t)
}

func check(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}

	if s.Val != t.Val {
		return false
	}

	return check(s.Left, t.Left) && check(s.Right, t.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
