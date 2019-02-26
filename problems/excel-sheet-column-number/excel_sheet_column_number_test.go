package excel_sheet_column_number

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestTitleToNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "A",
			expected: 1,
		},
		{
			input:    "C",
			expected: 3,
		},
		{
			input:    "Z",
			expected: 26,
		},
		{
			input:    "AA",
			expected: 27,
		},
		{
			input:    "AB",
			expected: 28,
		},
		{
			input:    "ZY",
			expected: 701,
		}, {
			input:    "AAA",
			expected: 703,
		},
		{
			input:    "XYZ",
			expected: 16900,
		},
	}
	for _, tc := range tests {
		output := titleToNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
