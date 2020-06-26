package main

/**
面试题 02.01. 移除重复节点

编写代码，移除未排序链表中的重复节点。保留最开始出现的节点。

示例1：
```
 输入：[1, 2, 3, 3, 2, 1]
 输出：[1, 2, 3]
```

示例2：
```
 输入：[1, 1, 1, 1, 2]
 输出：[1, 2]
```

提示：
- 链表长度在`[0, 20000]`范围内。
- 链表元素在`[0, 20000]`范围内。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveDuplicateNodes(head *ListNode) *ListNode {

	m := make(map[int]int)

	if head == nil {
		return head
	}
	curr := head
	m[curr.Val] = 1
	for curr.Next != nil {
		if _, ok := m[curr.Next.Val]; ok {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
		m[curr.Val] = 1
	}

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
