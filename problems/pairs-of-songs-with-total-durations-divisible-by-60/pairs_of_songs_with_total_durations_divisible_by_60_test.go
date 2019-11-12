package problem1010

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestNumPairsDivisibleBy60(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{30, 20, 150, 100, 40},
			expected: 3,
		},
		{
			input:    []int{60, 60, 60},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := numPairsDivisibleBy60(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
