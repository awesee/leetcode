package happy_number

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestIsHappy(t *testing.T) {
	tests := [...]caseType{
		{
			input:    19,
			expected: true,
		},
		{
			input:    1,
			expected: true,
		},
		{
			input:    0,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isHappy(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
