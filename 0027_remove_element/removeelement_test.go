package main

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	r := removeElement(nums, val)
	e := 2
	if r != e {
		t.Fatalf("Expected return to be %v, got %v", e, r)
	}

	rNums := nums[:e]
	eNums := []int{2, 2}
	slices.Sort(rNums)
	slices.Sort(eNums)

	if !slices.Equal(rNums, eNums) {
		t.Fatalf("Expected new slice to be %v, got %v", eNums, rNums)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	r := removeElement(nums, val)
	e := 5
	if r != e {
		t.Fatalf("Expected return to be %v, got %v", e, r)
	}

	rNums := nums[:e]
	eNums := []int{0, 1, 4, 0, 3}
	slices.Sort(rNums)
	slices.Sort(eNums)

	if !slices.Equal(rNums, eNums) {
		t.Fatalf("Expected new slice to be %v, got %v", eNums, rNums)
	}
}
