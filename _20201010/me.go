package main

/**
142. 环形链表 II

给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。
说明：不允许修改给定的链表。

示例1：
```
输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
```

示例2：
```
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
```

示例3：
```
输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。
```

进阶：
你是否可以不用额外空间解决此题？
/**
...
*/
// /**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */

// map
// func detectCycle(head *ListNode) *ListNode {

//     m := make(map[*ListNode]int)
//     for head != nil {
//         if _, ok := m[head] ; ok {
//             return head
//         }
//         m[head] = 1
//         head = head.Next
//     }
//     return nil
// }

// 快慢指针
func DetectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head.Next.Next, head
	var stack []*ListNode
	for fast != slow {
		stack = append(stack, slow)
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	p := head
	for p != slow {
		p = p.Next
		slow = slow.Next
	}
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
