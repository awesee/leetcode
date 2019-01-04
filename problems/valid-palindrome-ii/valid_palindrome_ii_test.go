package valid_palindrome_ii

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestValidPalindrome(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "aba",
			expected: true,
		},
		{
			input:    "abca",
			expected: true,
		},
		{
			input:    "hello",
			expected: false,
		},
		{
			input:    "abcdcbda",
			expected: true,
		},
		{
			input:    "abcbabbbca",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := validPalindrome(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
