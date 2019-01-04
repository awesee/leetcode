package count_binary_substrings

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestCountBinarySubstrings(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "00110011",
			expected: 6,
		},
		{
			input:    "10101",
			expected: 4,
		},
	}
	for _, tc := range tests {
		output := countBinarySubstrings(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
