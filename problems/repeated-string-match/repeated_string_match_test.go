package repeated_string_match

import "testing"

type caseType struct {
	a        string
	b        string
	expected int
}

func TestRepeatedStringMatch(t *testing.T) {
	tests := [...]caseType{
		{
			a:        "abcd",
			b:        "cdabcdab",
			expected: 3,
		},
		{
			a:        "abcdefg",
			b:        "cde",
			expected: 1,
		},
		{
			a:        "abc",
			b:        "abcabc",
			expected: 2,
		},
		{
			a:        "abc",
			b:        "abcd",
			expected: -1,
		},
	}
	for _, tc := range tests {
		output := repeatedStringMatch(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
