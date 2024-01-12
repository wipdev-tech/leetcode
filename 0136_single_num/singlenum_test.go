package singlenum

import "testing"

func TestAll(t *testing.T) {
	s1 := []int{2, 2, 1}
	e1 := 1
	r1 := singleNumber(s1)
	if e1 != r1 {
		t.Fatalf("Expected %v; got %v", e1, r1)
	}

	s2 := []int{4, 1, 2, 1, 2}
	e2 := 4
	r2 := singleNumber(s2)
	if e2 != r2 {
		t.Fatalf("Expected %v; got %v", e2, r2)
	}

	s3 := []int{1}
	e3 := 1
	r3 := singleNumber(s3)
	if e3 != r3 {
		t.Fatalf("Expected %v; got %v", e3, r3)
	}
}
