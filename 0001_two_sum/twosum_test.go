package twosum

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	expect := [][]int{{0, 1}, {1, 0}}

	for _, e := range expect {
		if fmt.Sprintf("%v", e) == fmt.Sprintf("%v", result) {
			return
		}
	}

	t.Fatalf("Expected one of %v; got %v", expect, result)
}

func TestCase2(t *testing.T) {
	result := twoSum([]int{3, 2, 4}, 6)
	expect := [][]int{{1, 2}, {2, 1}}

	for _, e := range expect {
		if fmt.Sprintf("%v", e) == fmt.Sprintf("%v", result) {
			return
		}
	}

	t.Fatalf("Expected one of %v; got %v", expect, result)
}

func TestCase3(t *testing.T) {
	result := twoSum([]int{3, 3}, 6)
	expect := [][]int{{0, 1}, {1, 0}}

	for _, e := range expect {
		if fmt.Sprintf("%v", e) == fmt.Sprintf("%v", result) {
			return
		}
	}

	t.Fatalf("Expected one of %v; got %v", expect, result)
}
