package singlenum

import "slices"

func singleNumber(nums []int) int {
	slices.Sort(nums)

	for i := 0; i < len(nums)-1; i++ {
		n := nums[i]
		if n != nums[i+1] {
			return n
		}
		i++
	}

	return nums[len(nums)-1]
}
