package containsdup

import "testing"

func TestCases(t *testing.T) {
	if !containsDuplicate([]int{1, 2, 3, 1}) {
		t.Fatalf("[1,2,3,1] should return true")
	}

	if containsDuplicate([]int{1, 2, 3, 4}) {
		t.Fatalf("[1,2,3,1] should return false")
	}

	if !containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) {
		t.Fatalf("[1,1,1,3,3,4,3,2,4,2] should return true")
	}
}
