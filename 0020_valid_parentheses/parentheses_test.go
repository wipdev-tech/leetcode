package parentheses

import "testing"

func Test(t *testing.T) {
	if !isValid("()") {
		t.Fatalf("`isValid(\"()\")` should be true")
	}
	if !isValid("()[]{}") {
		t.Fatalf("`isValid(\"()[]{}\")` should be true")
	}
	if isValid("(]") {
		t.Fatalf("`isValid(\"(]\")` should be false")
	}
}
