package verifying_an_alien_dictionary

import "testing"

type caseType struct {
	words    []string
	order    string
	expected bool
}

func TestIsAlienSorted(t *testing.T) {
	tests := [...]caseType{
		{
			words:    []string{"hello", "leetcode"},
			order:    "hlabcdefgijkmnopqrstuvwxyz",
			expected: true,
		},
		{
			words:    []string{"word", "world", "row"},
			order:    "worldabcefghijkmnpqstuvxyz",
			expected: false,
		},
		{
			words:    []string{"apple", "app"},
			order:    "abcdefghijklmnopqrstuvwxyz",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isAlienSorted(tc.words, tc.order)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.words, output, tc.expected)
		}
	}
}
