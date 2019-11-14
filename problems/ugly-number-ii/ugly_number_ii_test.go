package problem264

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestNthUglyNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    10,
			expected: 12,
		},
		{
			input:    9,
			expected: 10,
		},
		{
			input:    60,
			expected: 384,
		},
		{
			input:    90,
			expected: 1152,
		},
		{
			input:    1,
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := nthUglyNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
