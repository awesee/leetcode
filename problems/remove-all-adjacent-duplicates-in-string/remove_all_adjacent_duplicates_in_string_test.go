package remove_all_adjacent_duplicates_in_string

import "testing"

type caseType struct {
	input    string
	expected string
}

func TestRemoveDuplicates(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "abbaca",
			expected: "ca",
		},
		{
			input:    "aaabaca",
			expected: "abaca",
		},
	}
	for _, tc := range tests {
		output := removeDuplicates(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
