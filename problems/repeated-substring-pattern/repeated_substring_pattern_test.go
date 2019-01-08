package repeated_substring_pattern

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestRepeatedSubstringPattern(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "a",
			expected: false,
		},
		{
			input:    "abab",
			expected: true,
		},
		{
			input:    "aba",
			expected: false,
		},
		{
			input:    "abaaba",
			expected: true,
		},
		{
			input:    "abcabcabcabc",
			expected: true,
		},
		{
			input:    "abbaabbaabba",
			expected: true,
		},
		{
			input:    "abcabcabcdabcabcabcdabcabcabcd",
			expected: true,
		},
		{
			input:    "abaabaabac",
			expected: false,
		}, {
			input:    "babbabbabbabbab",
			expected: true,
		},
	}
	for _, tc := range tests {
		output := repeatedSubstringPattern(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
