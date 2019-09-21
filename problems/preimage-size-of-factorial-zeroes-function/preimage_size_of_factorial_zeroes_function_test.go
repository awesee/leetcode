package problem_793

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestPreimageSizeFZF(t *testing.T) {
	tests := [...]caseType{
		{
			input:    0,
			expected: 5,
		},
		{
			input:    5,
			expected: 0,
		},
		{
			input:    17,
			expected: 0,
		},
		{
			input:    11,
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := preimageSizeFZF(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
