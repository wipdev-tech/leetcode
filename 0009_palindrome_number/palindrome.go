package palindrome

import "strconv"

func isPalindrome(x int) bool {
	switch {
	case x < 0:
		return false

	case x == 0:
		return true

	default:
		xStr := strconv.Itoa(x)
		xRev := reverse(xStr)
		return xStr == xRev
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
