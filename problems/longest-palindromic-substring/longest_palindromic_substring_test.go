package longest_palindromic_substring

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestLongestPalindrome(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "babad",
			expected: "bab",
		},
		{
			input:    "cbbd",
			expected: "bb",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "a",
			expected: "a",
		},
		{
			input:    "abbba",
			expected: "abbba",
		},
	}
	for _, tc := range tests {
		output := longestPalindrome(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
