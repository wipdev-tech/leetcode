package stairs

import "testing"

func Test(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Fatalf("2 should return 2 but got %v", climbStairs(2))
	}

	if climbStairs(3) != 3 {
		t.Fatalf("3 should return 3 but got %v", climbStairs(3))
	}
	if climbStairs(4) != 5 {
		t.Fatalf("4 should return 5 but got %v", climbStairs(4))
	}
}
