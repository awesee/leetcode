package valid_boomerang

import "testing"

type caseType struct {
	input    [][]int
	expected bool
}

func TestIsBoomerang(t *testing.T) {
	tests := [...]caseType{
		{
			input: [][]int{
				{1, 1},
				{2, 3},
				{3, 2},
			},
			expected: true,
		},
		{
			input: [][]int{
				{1, 1},
				{2, 2},
				{3, 3},
			},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isBoomerang(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
