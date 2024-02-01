package plusone

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for {
		if i == -1 {
			digits = append([]int{1}, digits...)
			break
		}

		if digits[i] == 9 {
			digits[i] = 0
			i--
			continue
		}

		digits[i]++
		break
	}

	return digits
}
