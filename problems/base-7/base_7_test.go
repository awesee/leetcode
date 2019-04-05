package base_7

import "testing"

type caseType struct {
	input    int
	expected string
}

func TestConvertToBase7(t *testing.T) {
	tests := [...]caseType{
		{
			input:    100,
			expected: "202",
		},
		{
			input:    -7,
			expected: "-10",
		},
		{
			input:    -1,
			expected: "-1",
		},
		{
			input:    0,
			expected: "0",
		},
	}
	for _, tc := range tests {
		output := convertToBase7(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
