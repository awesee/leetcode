package factorial_trailing_zeroes

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestTrailingZeroes(t *testing.T) {
	tests := [...]caseType{
		{
			input:    3,
			expected: 0,
		},
		{
			input:    5,
			expected: 1,
		},
		{
			input:    7,
			expected: 1,
		},
		{
			input:    8,
			expected: 1,
		},
		{
			input:    10,
			expected: 2,
		},
		{
			input:    12,
			expected: 2,
		},
		{
			input:    25,
			expected: 6,
		},
	}
	for _, tc := range tests {
		output := trailingZeroes(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
