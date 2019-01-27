package string_without_aaa_or_bbb

import "testing"

type caseType struct {
	a        int
	b        int
	expected string
}

func TestStrWithout3a3b(t *testing.T) {
	tests := [...]caseType{
		{
			a:        6,
			b:        2,
			expected: "aabaabaa",
		},
		{
			a:        2,
			b:        6,
			expected: "bbabbabb",
		},
		{
			a:        1,
			b:        2,
			expected: "bba",
		},
		{
			a:        4,
			b:        1,
			expected: "aabaa",
		},
	}
	for _, tc := range tests {
		output := strWithout3a3b(tc.a, tc.b)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.a, tc.b, output, tc.expected)
		}
	}
}
