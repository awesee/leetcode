package reverse_words_in_a_string_iii

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestReverseWords(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "Let's take LeetCode contest",
			expected: "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tc := range tests {
		output := reverseWords(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
