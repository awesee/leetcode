package detect_capital

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestDetectCapitalUse(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "USA",
			expected: true,
		},
		{
			input:    "leetcode",
			expected: true,
		},
		{
			input:    "Google",
			expected: true,
		},
		{
			input:    "FlaG",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := detectCapitalUse(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
