package divisor_game

import "testing"

type caseType struct {
	input    int
	expected bool
}

func TestDivisorGame(t *testing.T) {
	tests := [...]caseType{
		{
			input:    2,
			expected: true,
		},
		{
			input:    3,
			expected: false,
		},
	}
	for _, tc := range tests {
		output := divisorGame(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
