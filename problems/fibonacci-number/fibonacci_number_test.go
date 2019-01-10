package fibonacci_number

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestFib(t *testing.T) {
	tests := [...]caseType{
		{
			input:    2,
			expected: 1,
		},
		{
			input:    3,
			expected: 2,
		},
		{
			input:    4,
			expected: 3,
		},
		{
			input:    1,
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := fib(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
