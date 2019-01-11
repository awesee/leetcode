package max_consecutive_ones

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			input:    []int{1, 0, 1, 1, 0, 1},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := findMaxConsecutiveOnes(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
