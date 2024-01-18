package dedupesorted

func removeDuplicates(nums []int) int {
	iCur := 0
	nCur := 0
	for i, n := range nums {
		if i == 0 || nCur != n {
			nCur = n
			nums[iCur], nums[i] = nums[i], nums[iCur]
			iCur++
		}
	}
	return iCur
}
