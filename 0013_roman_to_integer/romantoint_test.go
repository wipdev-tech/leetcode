package romantoint

import (
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	s1 := "III"
	e1 := 3
	r1 := romanToInt(s1)
	if e1 != r1 {
		log.Fatalf("Expected %v; got %v", e1, r1)
		// Explanation: III = 3.
	}

	s2 := "LVIII"
	e2 := 58
	r2 := romanToInt(s2)
	if e2 != r2 {
		log.Fatalf("Expected %v; got %v", e2, r2)
		// Explanation: L = 50, V= 5, III = 3.
	}

	s3 := "MCMXCIV"
	e3 := 1994
	r3 := romanToInt(s3)
	if e3 != r3 {
		log.Fatalf("Expected %v; got %v", e3, r3)
		// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
	}
}
