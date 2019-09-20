package problem_413

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 3, 5, 7, 9},
			expected: 6,
		},
		{
			input:    []int{7, 7, 7, 7},
			expected: 3,
		},
		{
			input:    []int{3, -1, -5, -9},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3, 4, 6, 8, 10},
			expected: 6,
		},
		{
			input:    []int{1, 2},
			expected: 0,
		},
		{
			input:    []int{1, 2, 4, 5, 7},
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := numberOfArithmeticSlices(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
