package merge

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	cur1 := list1
	cur2 := list2
	var merged *ListNode

	for cur1 != nil || cur2 != nil {
		switch {
		case cur1 == nil:
			return appendList(merged, cur2)

		case cur2 == nil:
			return appendList(merged, cur1)

		case cur1.Val < cur2.Val:
			merged = appendList(merged, &ListNode{Val: cur1.Val, Next: nil})
			cur1 = cur1.Next

		default:
			merged = appendList(merged, &ListNode{Val: cur2.Val, Next: nil})
			cur2 = cur2.Next
		}
	}

	return merged
}

func appendList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	lastNode := list1
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}

	lastNode.Next = list2
	return list1
}
