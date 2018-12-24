package valid_palindrome

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			input:    "",
			expected: true,
		},
		{
			input:    ".",
			expected: true,
		},
		{
			input:    "12321",
			expected: true,
		},
		{
			input:    "race a car",
			expected: false,
		},
		{
			input:    "hello, world",
			expected: false,
		},
	}

	for _, tc := range tests {
		output := isPalindrome(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
