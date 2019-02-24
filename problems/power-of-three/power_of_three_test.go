package power_of_three

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestIsPowerOfThree(t *testing.T) {
	tests := [...]caseType{
		{
			input:    27,
			expected: true,
		},
		{
			input:    0,
			expected: false,
		},
		{
			input:    1,
			expected: true,
		},
		{
			input:    9,
			expected: true,
		},
		{
			input:    45,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isPowerOfThree(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
