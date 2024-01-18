package dedupelist

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		return deleteDuplicates(head.Next)
	}

	return &ListNode{
		Val:  head.Val,
		Next: deleteDuplicates(head.Next),
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
