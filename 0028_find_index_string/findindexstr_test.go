package findindexstring

import "testing"

func TestCase1(t *testing.T) {
	r := strStr("sadbutsad", "sad")
	e := 0
	if r != e {
		t.Fatalf("Expected %v, found %v", e, r)
	}
}

func TestCase2(t *testing.T) {
	r := strStr("leetcode", "leeto")
	e := -1
	if r != e {
		t.Fatalf("Expected %v, found %v", e, r)
	}
}
