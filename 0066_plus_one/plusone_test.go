package plusone

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"Example 1", []int{1, 2, 3}, []int{1, 2, 4}},
		{"Example 2", []int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{"Example 3", []int{9}, []int{1, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := plusOne(tt.input)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, want %v", got, tt.output)
			}
		})
	}
}
