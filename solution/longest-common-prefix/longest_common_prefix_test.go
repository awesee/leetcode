package longest_common_prefix

import "testing"

type caseType struct {
	input    []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			input:    []string{"dog", "race", "car"},
			expected: "",
		},
		{
			input:    nil,
			expected: "",
		},
	}

	for _, tc := range tests {
		output := longestCommonPrefix(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
