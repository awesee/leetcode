package problem_1163

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestLastSubstring(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "abab",
			expected: "bab",
		},
		{
			input:    "abcdabc",
			expected: "dabc",
		},
		{
			input:    "ffbcdfeda",
			expected: "ffbcdfeda",
		},
		{
			input:    "hksljkjsfjv",
			expected: "v",
		},
	}
	for _, tc := range tests {
		output := lastSubstring(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
