package regular_expression_matching

import "testing"

type caseType struct {
	s        string
	p        string
	expected bool
}

func TestIsMatch(t *testing.T) {
	tests := [...]caseType{
		{
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			s:        "",
			p:        ".",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isMatch(tc.s, tc.p)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.s, tc.p, output, tc.expected)
		}
	}
}
