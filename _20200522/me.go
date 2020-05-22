package main

import "fmt"

/**
从前序与中序遍历序列构造二叉树

根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出
```
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
```

返回如下的二叉树：
```
    3
   / \
  9  20
    /  \
   15   7
```
*/

/**
嘿嘿...
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTreeWithMe(preorder []int, inorder []int) *TreeNode {
	var idxMap = map[int]int{}

	idxMap[1] = 10

	fmt.Println(idxMap[1])
	fmt.Println(idxMap[0])
	return nil

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
