package squares_of_a_sorted_array

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    []int
	expected []int
}

func TestSortedSquares(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		{
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tc := range tests {
		output := sortedSquares(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
