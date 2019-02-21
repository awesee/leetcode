package rotting_oranges

import "testing"

type caseType struct {
	input    [][]int
	expected int
}

func TestOrangesRotting(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			input: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			input: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			input: [][]int{
				{1, 2},
			},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := orangesRotting(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
