package bstdepth

import "testing"

func TestCase1(t *testing.T) {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	expect := 3
	result := maxDepth(tree)

	if result != expect {
		t.Fatalf("Expected %d; got %d", expect, result)
	}
}

func TestCase2(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	expect := 2
	result := maxDepth(tree)

	if result != expect {
		t.Fatalf("Expected %d; got %d", expect, result)
	}
}
