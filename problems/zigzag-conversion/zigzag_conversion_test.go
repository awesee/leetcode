package zigzag_conversion

import "testing"

type caseType struct {
	input    string
	numRows  int
	expected string
}

func TestConvert(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			input:    "AB",
			numRows:  1,
			expected: "AB",
		},
	}
	for _, tc := range tests {
		output := convert(tc.input, tc.numRows)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
