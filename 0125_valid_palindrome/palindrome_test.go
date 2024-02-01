package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	// Example 1
	if !isPalindrome("A man, a plan, a canal: Panama") {
		t.Errorf("isPalindrome(\"A man, a plan, a canal: Panama\") = false; want true")
	}

	// Example 2
	if isPalindrome("race a car") {
		t.Errorf("isPalindrome(\"race a car\") = true; want false")
	}

	// Example 3
	if !isPalindrome(" ") {
		t.Errorf("isPalindrome(\" \") = false; want true")
	}
}
