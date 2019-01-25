package find_pivot_index

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestPivotIndex(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 7, 3, 6, 5, 6},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3},
			expected: -1,
		},
		{
			input:    []int{0},
			expected: 0,
		},
		{
			input:    []int{0, 0},
			expected: 0,
		},
		{
			input:    []int{1, 2, 1},
			expected: 1,
		},
		{
			input:    []int{},
			expected: -1,
		},
	}
	for _, tc := range tests {
		output := pivotIndex(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
