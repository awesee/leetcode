package reverse_vowels_of_a_string

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestReverseVowels(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "hello",
			expected: "holle",
		},
		{
			input:    "leetcode",
			expected: "leotcede",
		},
	}
	for _, tc := range tests {
		output := reverseVowels(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
