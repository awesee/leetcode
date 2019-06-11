package container_with_most_water

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMaxArea(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			input:    []int{1, 8, 6, 30, 20, 6, 9, 10, 1},
			expected: 48,
		},
	}
	for _, tc := range tests {
		output := maxArea(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
