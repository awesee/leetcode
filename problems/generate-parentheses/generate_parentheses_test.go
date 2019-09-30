package problem_22

import (
	"reflect"
	"testing"
)

type caseType struct {
	input    int
	expected []string
}

func TestGenerateParenthesis(t *testing.T) {
	tests := [...]caseType{
		{
			input: 3,
			expected: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			input: 2,
			expected: []string{
				"(())",
				"()()",
			},
		},
		{
			input:    1,
			expected: []string{"()"},
		},
		{
			input:    0,
			expected: []string{},
		},
		{
			input: 5,
			expected: []string{
				"((((()))))",
				"(((()())))",
				"(((())()))",
				"(((()))())",
				"(((())))()",
				"((()(())))",
				"((()()()))",
				"((()())())",
				"((()()))()",
				"((())(()))",
				"((())()())",
				"((())())()",
				"((()))(())",
				"((()))()()",
				"(()((())))",
				"(()(()()))",
				"(()(())())",
				"(()(()))()",
				"(()()(()))",
				"(()()()())",
				"(()()())()",
				"(()())(())",
				"(()())()()",
				"(())((()))",
				"(())(()())",
				"(())(())()",
				"(())()(())",
				"(())()()()",
				"()(((())))",
				"()((()()))",
				"()((())())",
				"()((()))()",
				"()(()(()))",
				"()(()()())",
				"()(()())()",
				"()(())(())",
				"()(())()()",
				"()()((()))",
				"()()(()())",
				"()()(())()",
				"()()()(())",
				"()()()()()",
			},
		},
	}
	for _, tc := range tests {
		output := generateParenthesis(tc.input)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
