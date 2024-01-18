package dedupesorted

import (
	"slices"
	"testing"
)

func TestCase1(t *testing.T) {
	nums := []int{1, 1, 2}
	eNums := []int{1, 2} // expected sorted array
	eCap := len(nums)    // capacity of nums expected not to change
	eReturn := 2         // expected value returned from the func
	r := removeDuplicates(nums)

	if cap(nums) != eCap {
		t.Fatalf("Capacity changed from %v to %v", eCap, cap(nums))
	}

	if eReturn != r {
		t.Fatalf("Returned value expected to be %v; got %v", eReturn, r)
	}

	if !slices.Equal(nums[0:len(eNums)], eNums) {
		t.Fatalf("Sorted array expected to be %v; got %v", eNums, nums)
	}
}

func TestCase2(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	eNums := []int{0, 1, 2, 3, 4}
	eCap := len(nums)
	eReturn := 5
	r := removeDuplicates(nums)

	if cap(nums) != eCap {
		t.Fatalf("Capacity changed from %v to %v", eCap, cap(nums))
	}

	if eReturn != r {
		t.Fatalf("Returned value expected to be %v; got %v", eReturn, r)
	}

	if !slices.Equal(nums[0:len(eNums)], eNums) {
		t.Fatalf("Sorted array expected to be %v; got %v", eNums, nums)
	}
}
