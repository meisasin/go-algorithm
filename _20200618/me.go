package main

import (
	"strconv"
)

/**
我们从二叉树的根节点 root 开始进行深度优先搜索。

在遍历中的每个节点处，我们输出 D 条短划线（其中 D 是该节点的深度），然后输出该节点的值。（如果节点的深度为 D，则其直接子节点的深度为 D + 1。根节点的深度为 0）。
如果节点只有一个子节点，那么保证该子节点为左子节点。
给出遍历输出 S，还原树并返回其根节点 root。

示例1：
```
输入："1-2--3--4-5--6--7"
输出：[1,2,5,3,4,6,7]
```

示例2：
```
输入："1-2--3---4-5--6---7"
输出：[1,2,5,3,null,6,null,4,null,7]
```

示例2：
```
输入："1-401--349---90--88"
输出：[1,401,null,349,88,90]
```

提示：
- `原始树中的节点数介于 1 和 1000 之间。`
- `每个节点的值介于 1 和 10 ^ 9 之间。`
*/

/**
哇，自己完成了困难题唉，不容易啊。 不都发现了，通常节点树的题都可以用栈来解决
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func recoverFromPreorderWithMe(S string) *TreeNode {

	con, text := context(S)
	S = text

	val, _ := strconv.Atoi(con)
	root := &TreeNode{Val: val}

	next := root

	step, count := 0, 0
	stack := []*TreeNode{root}

	for {
		count, S = pre(S)
		con, S = context(S)
		if count == 0 || len(con) == 0 {
			break
		}

		val, _ := strconv.Atoi(con)
		curr := &TreeNode{Val: val}
		if count > step {
			next.Left = curr
			next = next.Left

			stack = append(stack, next)
			step++
		} else {
			for step >= count {
				stack = stack[0 : len(stack)-1]
				step--
			}
			next = stack[len(stack)-1]
			next.Right = curr
			next = next.Right

			stack = append(stack, next)
			step++
		}
	}

	return root
}

func pre(S string) (int, string) {
	count := 0
	for count < len(S) && S[count] == '-' {
		count++
	}
	return count, S[count:]
}

func context(S string) (string, string) {
	count := 0
	for count < len(S) && S[count] != '-' {
		count++
	}
	return S[0:count], S[count:]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
