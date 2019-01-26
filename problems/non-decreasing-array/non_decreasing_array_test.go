package non_decreasing_array

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestCheckPossibility(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 2, 5},
			expected: true,
		},
		{
			input:    []int{3, 2, 1},
			expected: false,
		},
		{
			input:    []int{3, 4, 2, 3},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 7, 5, 6},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 7, 1, 9},
			expected: true,
		},
	}
	for _, tc := range tests {
		output := checkPossibility(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
