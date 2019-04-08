package remove_outermost_parentheses

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestRemoveOuterParentheses(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "(()())(())",
			expected: "()()()",
		},
		{
			input:    "(()())(())(()(()))",
			expected: "()()()()(())",
		},
		{
			input:    "()()",
			expected: "",
		},
	}
	for _, tc := range tests {
		output := removeOuterParentheses(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
