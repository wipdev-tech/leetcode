package main

func removeElement(nums []int, val int) int {
	removed := make([]int, 0, len(nums))
	kept := make([]int, 0, len(nums))

	for _, v := range nums {
		if v == val {
			removed = append(removed, v)
		} else {
			kept = append(kept, v)
		}
	}

	for i, v := range kept {
		nums[i] = v
	}

	for i, v := range removed {
		nums[i+len(kept)] = v
	}

	return len(kept)
}
