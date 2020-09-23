package main

/**
617. 合并二叉树

给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
你需要将他们合并为一个新的二叉树。
合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点。


示例 1：
```
输入:
	Tree 1                     Tree 2
          1                         2
         / \                       / \
        3   2                     1   3
       /                           \   \
      5                             4   7
输出:
合并后的树:
	     3
	    / \
	   4   5
	  / \   \
	 5   4   7
```

注意: 合并必须从两个树的根节点开始。
*/

/**
... 有点蠢，既然都准备用新的了，左右如果一边没有，直接返回就可以了
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func MergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {

	if t1 == nil && t2 == nil {
		return nil
	}
	ans := 0
	var lt, rt *TreeNode
	if t1 != nil && t2 != nil {
		ans += t1.Val
		ans += t2.Val
		lt, rt = MergeTrees(t1.Left, t2.Left), MergeTrees(t1.Right, t2.Right)
	} else if t1 != nil {
		ans += t1.Val
		lt, rt = MergeTrees(t1.Left, nil), MergeTrees(t1.Right, nil)
	} else {
		ans += t2.Val
		lt, rt = MergeTrees(nil, t2.Left), MergeTrees(nil, t2.Right)
	}

	root := &TreeNode{}
	root.Val = ans
	root.Left = lt
	root.Right = rt
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}