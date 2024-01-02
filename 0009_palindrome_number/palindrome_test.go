package palindrome

import "testing"

func Test(t *testing.T) {
	if !isPalindrome(121) {
		t.Fatalf("`isPalindrome(121)` should be true")
	}

	if isPalindrome(-121) {
		t.Fatalf("`isPalindrome(-121)` should be false")
	}

	if isPalindrome(10) {
		t.Fatalf("`isPalindrome(10)` should be false")
	}
}
