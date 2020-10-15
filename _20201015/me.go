package main

/**
116. 填充每个节点的下一个右侧节点指针

给定一个完美二叉树，其所有叶子节点都在同一层，每个父节点都有两个子节点。
二叉树定义如下：

```
struct Node {
  int val;
  Node *left;
  Node *right;
  Node *next;
}
```
填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

初始状态下，所有 next 指针都被设置为 NULL。

示例：
![_1.png](./source/_1.png)
```
输入：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":null,"right":null,"val":4},"next":null,"right":{"$id":"4","left":null,"next":null,"right":null,"val":5},"val":2},"next":null,"right":{"$id":"5","left":{"$id":"6","left":null,"next":null,"right":null,"val":6},"next":null,"right":{"$id":"7","left":null,"next":null,"right":null,"val":7},"val":3},"val":1}
输出：{"$id":"1","left":{"$id":"2","left":{"$id":"3","left":null,"next":{"$id":"4","left":null,"next":{"$id":"5","left":null,"next":{"$id":"6","left":null,"next":null,"right":null,"val":7},"right":null,"val":6},"right":null,"val":5},"right":null,"val":4},"next":{"$id":"7","left":{"$ref":"5"},"next":null,"right":{"$ref":"6"},"val":3},"right":{"$ref":"4"},"val":2},"next":null,"right":{"$ref":"7"},"val":1}
解释：给定二叉树如图 A 所示，你的函数应该填充它的每个 next 指针，以指向其下一个右侧节点，如图 B 所示。
```

提示：
- 你只能使用常量级额外空间。
- 使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。
*/

// ------------------------------------------------------------------------------------------

/**
...
*/

// ------------------------------------------------------------------------------------------

// /**
//  * Definition for a Node.
//  * type Node struct {
//  *     Val int
//  *     Left *Node
//  *     Right *Node
//  *     Next *Node
//  * }
//  */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func Connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root.Left != nil {
			root.Left.Next = root.Right
			if root.Next != nil {
				root.Right.Next = root.Next.Left
			}
			dfs(root.Left)
			dfs(root.Right)
		}
	}
	dfs(root)

	return root
}

// -------------
// func connect(root *Node) *Node {
//     if root == nil {
//         return nil
//     }
//     var stack []*Node
//     stack = append(stack, root)

//     for len(stack) > 0 {
//         sl := len(stack)
//         for i := 0 ; i < sl; i ++ {
//             curr := stack[i]
//             if curr.Left != nil {
//                 stack = append(stack, curr.Left)
//             }
//             if curr.Right != nil {
//                 stack = append(stack, curr.Right)
//             }
//             if i != sl - 1 {
//                 curr.Next = stack[i + 1]
//             }
//         }
//         stack = stack[sl:]
//     }

//     return root
// }

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
