package move_zeroes

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			input:    []int{0, 1, 0, 2, 3, 3, 0, 7, 8},
			expected: []int{1, 2, 3, 3, 7, 8, 0, 0, 0},
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		output := make([]int, len(tc.input))
		copy(output, tc.input)
		moveZeroes(output)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
