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

	for _, test := range tests {
		output := longestCommonPrefix(test.input)
		if output != test.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", test.input, output, test.expected)
		}
	}
}
