package most_common_word

import "testing"

type caseType struct {
	paragraph string
	banned    []string
	expected  string
}

func TestMostCommonWord(t *testing.T) {
	tests := [...]caseType{
		{
			paragraph: "Bob hit a ball, the hit BALL flew far after it was hit.",
			banned:    []string{"hit"},
			expected:  "ball",
		},
		{
			paragraph: "a, a, a, a, b,b,b,c, c",
			banned:    []string{"a"},
			expected:  "b",
		},
		{
			paragraph: "Bob",
			banned:    []string{},
			expected:  "bob",
		},
	}
	for _, tc := range tests {
		output := mostCommonWord(tc.paragraph, tc.banned)
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.paragraph, tc.banned, output, tc.expected)
		}
	}
}
