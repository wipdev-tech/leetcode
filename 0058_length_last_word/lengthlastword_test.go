package lengthlastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{input: "Hello World", output: 5},
		{input: "   fly me   to   the moon  ", output: 4},
		{input: "luffy is still joyboy", output: 6},
	}

	for _, tc := range testCases {
		result := lengthOfLastWord(tc.input)
		if result != tc.output {
			t.Errorf("lengthOfLastWord(%q) = %d; want %d", tc.input, result, tc.output)
		}
	}
}
