package rotated_digits

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestRotatedDigits(t *testing.T) {
	tests := [...]caseType{
		{
			input:    10,
			expected: 4,
		},
		{
			input:    23,
			expected: 11,
		},
		{
			input:    100,
			expected: 40,
		},
		{
			input:    200,
			expected: 81,
		},
	}
	for _, tc := range tests {
		output := rotatedDigits(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
