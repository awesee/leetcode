package problem_565

import (
	"testing"
)

type caseType struct {
	input    []int
	expected int
}

func TestArrayNesting(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{5, 4, 0, 3, 1, 6, 2},
			expected: 4,
		},
		{
			input:    []int{0, 3, 1, 5, 4, 6, 2},
			expected: 5,
		},
		{
			input:    []int{6, 2, 5, 4, 0, 3, 1},
			expected: 7,
		},
		{
			input:    []int{3, 1, 5, 4, 0, 6, 2},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := arrayNesting(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
