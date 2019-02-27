package reverse_words_in_a_string

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestReverseWords(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			input:    "  hello world!  ",
			expected: "world! hello",
		},
		{
			input:    "a good   example",
			expected: "example good a",
		},
	}
	for _, tc := range tests {
		output := reverseWords(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
