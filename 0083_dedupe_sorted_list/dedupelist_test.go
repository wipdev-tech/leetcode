package dedupelist

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
	l := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				2,
				nil,
			},
		},
	}
	r := deleteDuplicates(l)
	if stringify(r) != "1 -> 2 -> nil" {
		t.Fatalf("Expected 1 -> 2 -> nil, got %v", stringify(r))
	}
}

func TestCase2(t *testing.T) {
	l1 := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				2,
				&ListNode{
					3,
					&ListNode{
						3,
						nil,
					},
				},
			},
		},
	}
	r1 := deleteDuplicates(l1)
	if stringify(r1) != "1 -> 2 -> 3 -> nil" {
		t.Fatalf("Expected 1 -> 2 -> 3 -> nil, got %v", stringify(r1))
	}
}

func TestCase3(t *testing.T) {
	l := &ListNode{
		1,
		&ListNode{
			1,
			&ListNode{
				1,
				nil,
			},
		},
	}
	r := deleteDuplicates(l)
	if stringify(r) != "1 -> nil" {
		t.Fatalf("Expected 1 -> nil, got %v", stringify(r))
	}
}
