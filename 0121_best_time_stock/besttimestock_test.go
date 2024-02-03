package besttimestock

import "testing"

func TestMaxProfit(t *testing.T) {
	type testCase struct {
		stocks   []int
		expected int
	}

	testCases := []testCase{
		testCase{[]int{7, 1, 5, 3, 6, 4}, 5},
		testCase{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, c := range testCases {
		if maxProfit(c.stocks) != c.expected {
			t.Fatalf("Expected %v, got %v", c.expected, maxProfit(c.stocks))
		}
	}
}
