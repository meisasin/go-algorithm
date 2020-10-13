package main

/**
24. 两两交换链表中的节点

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例1：
```
给定 1->2->3->4, 你应该返回 2->1->4->3.
```
*/

// ------------------------------------------------------------------------------------------

/**
写得我自己都稀里糊涂的
*/

// ------------------------------------------------------------------------------------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {

	ans := &ListNode{}
	tmp, atmp := head, ans
	for tmp != nil {
		if tmp.Next == nil {
			atmp.Next = tmp
			break
		} else {
			first, second, third := tmp.Next, tmp, tmp.Next.Next
			first.Next = nil
			second.Next = third
			first.Next = second
			atmp.Next = first

			tmp = third
			atmp = atmp.Next.Next
		}
	}

	return ans.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
