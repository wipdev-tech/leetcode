package searchinsertpos

import (
	"testing"
)

func TestCase1(t *testing.T) {
	r := searchInsert([]int{1, 3, 5, 6}, 5)
	e := 2
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase2(t *testing.T) {
	r := searchInsert([]int{1, 3, 5, 6}, 2)
	e := 1
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase3(t *testing.T) {
	r := searchInsert([]int{1, 3, 5, 6}, 7)
	e := 4
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase4(t *testing.T) {
	r := searchInsert([]int{1}, 1)
	e := 0
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase5(t *testing.T) {
	r := searchInsert([]int{1}, 2)
	e := 1
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase6(t *testing.T) {
	r := searchInsert([]int{1, 3}, 1)
	e := 0
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase7(t *testing.T) {
	r := searchInsert([]int{1, 3}, 2)
	e := 1
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase8(t *testing.T) {
	r := searchInsert([]int{1, 3, 5}, 4)
	e := 2
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}

func TestCase9(t *testing.T) {
	r := searchInsert([]int{3, 4, 5, 6, 7, 8}, 6)
	e := 3
	if r != e {
		t.Fatalf("Expected %v, got %v", e, r)
	}
}
