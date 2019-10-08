package problem_290

import "testing"

type caseType struct {
	pattern  string
	str      string
	expected bool
}

func TestWordPattern(t *testing.T) {
	tests := [...]caseType{
		{
			pattern:  "abba",
			str:      "dog cat cat dog",
			expected: true,
		},
		{
			pattern:  "abba",
			str:      "dog cat cat fish",
			expected: false,
		},
		{
			pattern:  "aaaa",
			str:      "dog cat cat dog",
			expected: false,
		},
		{
			pattern:  "abba",
			str:      "dog dog dog dog",
			expected: false,
		},
		{
			pattern:  "abba",
			str:      "dog cat cat dog fish",
			expected: false,
		},
		{
			pattern:  "abba",
			str:      "b a a b",
			expected: true,
		},
		{
			pattern:  "abba",
			str:      "b c c b",
			expected: true,
		},
	}
	for _, tc := range tests {
		output := wordPattern(tc.pattern, tc.str)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.pattern, output, tc.expected)
		}
	}
}
