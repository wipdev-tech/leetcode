package sametree

import "testing"

func TestCase1(t *testing.T) {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	q := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}

	if !isSameTree(p, q) {
		t.Fatalf("p and q should be the same tree")
	}
}

func TestCase2(t *testing.T) {
	p := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}
	q := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	if isSameTree(p, q) {
		t.Fatalf("p and q should NOT be the same tree")
	}
}

func TestCase3(t *testing.T) {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}
	q := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	if isSameTree(p, q) {
		t.Fatalf("p and q should NOT be the same tree")
	}
}
