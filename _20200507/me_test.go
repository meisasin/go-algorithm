package main

import (
	"fmt"
	"testing"
)

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
func TestIsSubtree(t *testing.T) {

	sRoot := TreeNode{Val: 3}

	sLeft1 := TreeNode{Val: 4}
	sLeft1_left := TreeNode{Val: 1}
	sLeft1_right := TreeNode{Val: 2}

	//sLeft1_right_left := TreeNode{ Val: 0}

	sRight1 := TreeNode{Val: 5}

	//sLeft1_right.Left = &sLeft1_right_left
	sLeft1.Left = &sLeft1_left
	sLeft1.Right = &sLeft1_right
	sRoot.Left = &sLeft1
	sRoot.Right = &sRight1

	tRoot := TreeNode{
		Val: 4,
	}

	tLeft1 := TreeNode{Val: 1}
	tRight1 := TreeNode{Val: 2}
	tRoot.Left = &tLeft1
	tRoot.Right = &tRight1

	fmt.Println("isSubtree: ", IsSubtree(&sRoot, &tRoot))

}
