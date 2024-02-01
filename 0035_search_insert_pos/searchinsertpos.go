package searchinsertpos

import (
	"fmt"
	"math"
)

// UNSOLVED
func searchInsert(nums []int, target int) int {
	fmt.Println(nums, target)
	if target > nums[len(nums)-1] {
		return len(nums)
	}

	if target < nums[0] {
		return 0
	}

	lo := 0
	hi := len(nums) - 1

	for lo < hi-1 {
		md := int(math.Floor(float64(hi) / 2.0))
		fmt.Println(lo, md, hi)
		if target == nums[md] {
			return md
		}

		if target > nums[md] {
			lo = md + 1
		}

		if target < nums[md] {
			hi = md - 1
		}
	}
	fmt.Println(lo, hi)

	switch {
	case target == nums[lo] || target < nums[lo]:
		return lo
	case target == nums[hi]:
		return hi
	case target > nums[hi]:
		return hi + 1
	case target > nums[lo] && target < nums[hi]:
		return hi
	default:
		return -1
	}
}
