package main

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 2, 3, 0, 0, 0}
	merge(nums, 3, []int{2, 5, 6}, 3)
	e := []int{1, 2, 2, 3, 5, 6}
	if !slices.Equal(nums, e) {
		t.Fatalf("Expected %v, got %v", e, nums)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{1}
	merge(nums, 1, []int{}, 0)
	e := []int{1}
	if !slices.Equal(nums, e) {
		t.Fatalf("Expected %v, got %v", e, nums)
	}
}

func TestCase3(t *testing.T) {
	nums := []int{0}
	merge(nums, 0, []int{1}, 1)
	e := []int{1}
	if !slices.Equal(nums, e) {
		t.Fatalf("Expected %v, got %v", e, nums)
	}
}
