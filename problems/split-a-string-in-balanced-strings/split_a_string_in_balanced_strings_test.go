package problem1221

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestBalancedStringSplit(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "RLRRLLRLRL",
			expected: 4,
		},
		{
			input:    "RLLLLRRRLR",
			expected: 3,
		},
		{
			input:    "LLLLRRRR",
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := balancedStringSplit(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
