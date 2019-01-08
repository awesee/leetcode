package reverse_string_ii

import "testing"

type caseType struct {
	s        string
	k        int
	expected string
}

func TestReverseStr(t *testing.T) {
	tests := [...]caseType{
		{
			s:        "abcdefg",
			k:        2,
			expected: "bacdfeg",
		},
		{
			s:        "abcdefg",
			k:        4,
			expected: "dcbaefg",
		},
		{
			s:        "abcd",
			k:        4,
			expected: "dcba",
		},
		{
			s:        "abcdefg",
			k:        8,
			expected: "gfedcba",
		},
	}
	for _, tc := range tests {
		output := reverseStr(tc.s, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.s, tc.k, output, tc.expected)
		}
	}
}
