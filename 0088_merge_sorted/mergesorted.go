package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1)
	i1 := 0 // current idx taking from nums1Copy
	i2 := 0 // current idx taking from nums2
	it := 0 // current idx receiving into nums1

	for i1 < m && i2 < n {
		if nums1Copy[i1] < nums2[i2] {
			nums1[it] = nums1Copy[i1]
			it++
			i1++
			continue
		}

		if nums1Copy[i1] > nums2[i2] {
			nums1[it] = nums2[i2]
			it++
			i2++
			continue
		}

		nums1[it] = nums1Copy[i1]
		nums1[it+1] = nums2[i2]
		it += 2
		i1++
		i2++
	}

	for i1 < m {
		nums1[it] = nums1Copy[i1]
		it++
		i1++
	}

	for i2 < n {
		nums1[it] = nums2[i2]
		it++
		i2++
	}
}
