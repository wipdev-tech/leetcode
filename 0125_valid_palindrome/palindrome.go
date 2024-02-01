package palindrome

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	lower := strings.ToLower(s)
	re := regexp.MustCompile("[^A-Za-z0-9]")
	clean := re.ReplaceAll([]byte(lower), []byte{})

	if string(clean) == "" {
		return true
	}

	for i := 0; i <= len(clean)/2; i++ {
		j := len(clean) - 1 - i
		if clean[i] != clean[j] {
			return false
		}
	}

	return true
}
