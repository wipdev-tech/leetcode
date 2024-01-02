package twosum

func twoSum(nums []int, target int) []int {
	var lookup = map[int]int{}

	for idx, val := range nums {
		needVal := target - val
		needIdx, lookupSuccess := lookup[needVal]

		if lookupSuccess && needIdx != idx {
			return []int{idx, needIdx}
		} else {
			lookup[val] = idx
		}
	}

	return []int{-1, -1}
}
