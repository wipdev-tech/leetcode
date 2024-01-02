package commonprefix

import "testing"

func TestCase1(t *testing.T) {
	case1 := []string{"flower", "flow", "flight"}
	if r := longestCommonPrefix(case1); r != "fl" {
		t.Fatalf("Expected 'fl'; got %v", r)
	}
}

func TestCase2(t *testing.T) {
	case2 := []string{"dog", "racecar", "car"}
	if r := longestCommonPrefix(case2); r != "" {
		t.Fatalf("Expected ''; got %v", r)
	}
}
