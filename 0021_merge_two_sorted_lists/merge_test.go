package merge

import (
	"fmt"
	"testing"
)

func stringify(list *ListNode) string {
	if list == nil {
		return "nil"
	}
	return fmt.Sprintf("%v -> ", list.Val) + stringify(list.Next)
}

func TestCase1(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	resultStr := stringify(mergeTwoLists(list1, list2))
	expectStr := "1 -> 1 -> 2 -> 3 -> 4 -> 4 -> nil"

	if resultStr != expectStr {
		t.Fatalf("Expected %s\nFound    %s", expectStr, resultStr)
	}
}

func TestCase2(t *testing.T) {
	var list1 *ListNode
	var list2 *ListNode

	resultStr := stringify(mergeTwoLists(list1, list2))
	expectStr := "nil"

	if resultStr != expectStr {
		t.Fatalf("Expected %s\nFound    %s", expectStr, resultStr)
	}
}

func TestCase3(t *testing.T) {
	var list1 *ListNode
	list2 := &ListNode{
		Val:  0,
		Next: nil,
	}

	resultStr := stringify(mergeTwoLists(list1, list2))
	expectStr := "0 -> nil"

	if resultStr != expectStr {
		t.Fatalf("Expected %s\nFound    %s", expectStr, resultStr)
	}
}
