package power_of_two

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestIsPowerOfTwo(t *testing.T) {
	tests := [...]caseType{
		{
			input:    0,
			expected: false,
		},
		{
			input:    1,
			expected: true,
		},
		{
			input:    12,
			expected: false,
		},
		{
			input:    16,
			expected: true,
		},
		{
			input:    218,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isPowerOfTwo(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
