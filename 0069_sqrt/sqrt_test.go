package sqrt

import (
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct {
		name   string
		input  int
		output int
	}{
		{"Example 1", 4, 2},
		{"Example 2", 8, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mySqrt(tt.input)
			if tt.output == tt.input {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		})
	}
}
