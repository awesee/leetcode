package sqrtx

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestMySqrt(t *testing.T) {
	tests := [...]caseType{
		{
			input:    4,
			expected: 2,
		},
		{
			input:    8,
			expected: 2,
		},
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := mySqrt(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
