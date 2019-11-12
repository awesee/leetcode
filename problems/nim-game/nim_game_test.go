package problem_292

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestCanWinNim(t *testing.T) {
	tests := [...]caseType{
		{
			input:    4,
			expected: false,
		},
		{
			input:    3,
			expected: true,
		},
	}
	for _, tc := range tests {
		output := canWinNim(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
