package problem_233

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestCountDigitOne(t *testing.T) {
	tests := [...]caseType{
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    12,
			expected: 5,
		},
		{
			input:    99,
			expected: 20,
		},
		{
			input:    100,
			expected: 21,
		},
		{
			input:    123,
			expected: 57,
		},
		{
			input:    1234,
			expected: 689,
		},
		{
			input:    12345,
			expected: 8121,
		},
		{
			input:    123456,
			expected: 93553,
		},
		{
			input:    1234567,
			expected: 1058985,
		},
		{
			input:    12345678,
			expected: 11824417,
		},
		{
			input:    123456789,
			expected: 130589849,
		},
	}
	for _, tc := range tests {
		output := countDigitOne(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
