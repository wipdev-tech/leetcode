package addtwonumbers

import "testing"

func TestListToNumber(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	result := listToNumber(l1)
	expect := 342
	if result != expect {
		t.Fatalf("Expected %v, got %v", expect, result)
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result = listToNumber(l2)
	expect = 465
	if result != expect {
		t.Fatalf("Expected %v, got %v", expect, result)
	}
}

func TestNumberToList(t *testing.T) {
	n1 := 342
	result := listToNumber(numberToList(n1))
	if result != n1 {
		t.Fatalf("Expected %v, got %v", n1, result)
	}

	n2 := 465
	result = listToNumber(numberToList(n2))
	if result != n2 {
		t.Fatalf("Expected %v, got %v", n2, result)
	}
}

func TestCase1(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	resultNum := listToNumber(addTwoNumbers(l1, l2))
	expectNum := 807

	if resultNum != expectNum {
		t.Fatalf("Expected %v; got %v", expectNum, resultNum)
	}
}

func TestCase2(t *testing.T) {
	l1 := &ListNode{
		Val: 0,
	}
	l2 := &ListNode{
		Val: 0,
	}

	resultNum := listToNumber(addTwoNumbers(l1, l2))
	expectNum := 0

	if resultNum != expectNum {
		t.Fatalf("Expected %v; got %v", expectNum, resultNum)
	}
}

func TestCase3(t *testing.T) {
	l1 := numberToList(9999999)
	l2 := numberToList(9999)

	resultNum := listToNumber(addTwoNumbers(l1, l2))
	expectNum := 10009998

	if resultNum != expectNum {
		t.Fatalf("Expected %v; got %v", expectNum, resultNum)
	}
}
