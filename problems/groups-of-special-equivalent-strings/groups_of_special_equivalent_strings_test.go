package groups_of_special_equivalent_strings

import "testing"

type caseType struct {
	input    []string
	expected int
}

func TestNumSpecialEquivGroups(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"a", "b", "c", "a", "c", "c"},
			expected: 3,
		},
		{
			input:    []string{"aa", "bb", "ab", "ba"},
			expected: 4,
		},
		{
			input:    []string{"abc", "acb", "bac", "bca", "cab", "cba"},
			expected: 3,
		},
		{
			input:    []string{"abcd", "cdab", "adcb", "cbad"},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := numSpecialEquivGroups(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
