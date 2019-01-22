package valid_mountain_array

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestValidMountainArray(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 1},
			expected: false,
		},
		{
			input:    []int{3, 5, 5},
			expected: false,
		},
		{
			input:    []int{0, 3, 2, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 2},
			expected: true,
		},
		{
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
			expected: false,
		},
		{
			input:    []int{1, 2, 3, 4, 5},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := validMountainArray(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
