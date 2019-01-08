package longest_uncommon_subsequence_i

import "testing"

type caseType struct {
	a        string
	b        string
	expected int
}

func TestFindLUSlength(t *testing.T) {
	tests := [...]caseType{
		{
			a:        "aba",
			b:        "cdc",
			expected: 3,
		},
		{
			a:        "aba",
			b:        "abcdef",
			expected: 6,
		},
		{
			a:        "abcdef",
			b:        "cdc",
			expected: 6,
		},
		{
			a:        "abc",
			b:        "abc",
			expected: -1,
		},
	}
	for _, tc := range tests {
		output := findLUSlength(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
