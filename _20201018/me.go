package main

/**
19. 删除链表的倒数第N个节点

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：
```
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5.
```

说明：
- 给定的 n 保证是有效的。

进阶：
- 你能尝试使用一趟扫描实现吗？
*/

// ------------------------------------------------------------------------------------------

/**
我这也算是一趟扫描了吧，不过我这是递归做的，好像也不太算
*/

// ------------------------------------------------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	var delNode func(head *ListNode) int
	delNode = func(head *ListNode) int {
		if head == nil {
			return 0
		}
		c := delNode(head.Next)
		if c == n {
			if head.Next.Next != nil {
				head.Next = head.Next.Next
			} else {
				head.Next = nil
			}
			n = -1
		}
		return c + 1
	}

	delNode(head)

	if n != -1 {
		return head.Next
	}

	return head
}

// -------------

// func removeNthFromEnd(head *ListNode, n int) *ListNode {

//     var stack []*ListNode
//     tmp := head
//     for tmp != nil {
//         stack = append(stack, tmp)
//         tmp = tmp.Next
//     }
//     for i := 0 ; i < n; i ++ {
//         stack = stack[:len(stack) - 1]
//     }
//     if len(stack) == 0 {
//         return head.Next
//     }
//     last := stack[len(stack) - 1]
//     if last.Next.Next == nil {
//         last.Next = nil
//     } else {
//         last.Next = last.Next.Next
//     }
//     return head
// }

// -------------

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     c := deleteNode(head, n)
//     if c == n {
//         return head.Next
//     }
//     return head
// }

// func deleteNode(head *ListNode, n int) int {
//     if head == nil {
//         return 0
//     }
//     c := deleteNode(head.Next, n)
//     if c == n {
//         head.Next = head.Next.Next
//     }
//     return c + 1
// }

type ListNode struct {
	Val  int
	Next *ListNode
}
