package toeplitz_matrix

import "testing"

type caseType struct {
	input    [][]int
	expected bool
}

func TestIsToeplitzMatrix(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			expected: true,
		},
		{
			input: [][]int{
				{1, 2},
				{2, 2},
			},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isToeplitzMatrix(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
