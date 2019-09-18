package problem_456

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestFind132pattern(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			input:    []int{3, 1, 4, 2},
			expected: true,
		},
		{
			input:    []int{-1, 3, 2, 0},
			expected: true,
		},
	}
	for _, tc := range tests {
		output := find132pattern(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
