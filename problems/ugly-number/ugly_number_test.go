package problem263

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestIsUgly(t *testing.T) {
	tests := [...]caseType{
		{
			input:    1,
			expected: true,
		},
		{
			input:    6,
			expected: true,
		},
		{
			input:    8,
			expected: true,
		},
		{
			input:    14,
			expected: false,
		},
		{
			input:    0,
			expected: false,
		},
		{
			input:    -30,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isUgly(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
