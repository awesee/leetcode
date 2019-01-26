package magic_squares_in_grid

import "testing"

type caseType struct {
	input    [][]int
	expected int
}

func TestNumMagicSquaresInside(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{4, 3, 8, 4},
				{9, 5, 1, 9},
				{2, 7, 6, 2},
			},
			expected: 1,
		},
		{
			input: [][]int{
				{5, 5, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
			expected: 0,
		},
		{
			input: [][]int{
				{1, 8, 6},
				{10, 5, 0},
				{4, 2, 9},
			},
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := numMagicSquaresInside(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
