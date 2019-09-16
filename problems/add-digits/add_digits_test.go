package add_digits

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestAddDigits(t *testing.T) {
	tests := [...]caseType{
		{
			input:    38,
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := addDigits(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
