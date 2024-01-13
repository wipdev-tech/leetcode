package threesum

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	if nums[0] > 0 || nums[len(nums)-1] < 0 {
		return [][]int{}
	}

	if nums[0] == nums[len(nums)-1] {
		if nums[0] == 0 {
			return [][]int{{0, 0, 0}}
		} else {
			return [][]int{}
		}
	}

	triplets := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if k < 0 {
					continue
				}

				if nums[i]+nums[j]+nums[k] == 0 {
					newTriplet := []int{nums[i], nums[j], nums[k]}
					slices.Sort(newTriplet)

					new := true
					for _, t := range triplets {
						if slices.Equal(t, newTriplet) {
							new = false
						}
					}

					if new {
						triplets = append(triplets, newTriplet)
					}
				}
			}
		}
	}

	return triplets
}
