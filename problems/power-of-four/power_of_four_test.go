package power_of_four

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestIsPowerOfFour(t *testing.T) {
	tests := [...]caseType{
		{
			input:    16,
			expected: true,
		},
		{
			input:    5,
			expected: false,
		},
		{
			input:    6,
			expected: false,
		},
		{
			input:    12,
			expected: false,
		},
		{
			input:    0,
			expected: false,
		},
		{
			input:    1,
			expected: true,
		},
	}
	for _, tc := range tests {
		output := isPowerOfFour(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
