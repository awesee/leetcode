package height_checker

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestHeightChecker(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 1, 4, 2, 1, 3},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := heightChecker(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
