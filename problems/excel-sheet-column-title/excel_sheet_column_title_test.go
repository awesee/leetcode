package problem_168

import "testing"

type caseType struct {
	input    int
	expected string
}

func TestConvertToTitle(t *testing.T) {
	tests := [...]caseType{
		{
			input:    1,
			expected: "A",
		},
		{
			input:    3,
			expected: "C",
		},
		{
			input:    26,
			expected: "Z",
		},
		{
			input:    52,
			expected: "AZ",
		},
		{
			input:    27,
			expected: "AA",
		},
		{
			input:    28,
			expected: "AB",
		},
		{
			input:    701,
			expected: "ZY",
		},
		{
			input:    703,
			expected: "AAA",
		},
		{
			input:    16900,
			expected: "XYZ",
		},
	}
	// Solution 1
	for _, tc := range tests {
		output := convertToTitle(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
	// Solution 2
	for _, tc := range tests {
		output := convertToTitle2(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
