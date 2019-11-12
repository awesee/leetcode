package problem_877

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestStoneGame(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{5, 3, 4, 5},
			expected: true,
		},
		{
			input:    []int{2, 5, 7, 3},
			expected: true,
		},
	}
	for _, tc := range tests {
		output := stoneGame(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
