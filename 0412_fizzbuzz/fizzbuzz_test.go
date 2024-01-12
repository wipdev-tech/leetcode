package fizzbuzz

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	n1 := 3
	e1 := []string{"1", "2", "Fizz"}
	r1 := fizzBuzz(n1)

	if !slices.Equal(e1, r1) {
		t.Fatalf("Expected %v, found %v", e1, r1)
	}

	n2 := 5
	e2 := []string{"1", "2", "Fizz", "4", "Buzz"}
	r2 := fizzBuzz(n2)

	if !slices.Equal(e2, r2) {
		t.Fatalf("Expected %v, found %v", e2, r2)
	}

	n3 := 15
	e3 := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	r3 := fizzBuzz(n3)

	if !slices.Equal(e3, r3) {
		t.Fatalf("Expected %v, found %v", e3, r3)
	}
}
