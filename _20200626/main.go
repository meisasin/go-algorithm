package main

func main() {

}

func removeDuplicateNodes(head *ListNode) *ListNode {
	ob := head
	for ob != nil {
		oc := ob
		for oc.Next != nil {
			if oc.Next.Val == ob.Val {
				oc.Next = oc.Next.Next
			} else {
				oc = oc.Next
			}
		}
		ob = ob.Next
	}
	return head
}
