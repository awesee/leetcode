package problem_717

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestIsOneBitCharacter(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 0, 0},
			expected: true,
		},
		{
			input:    []int{1, 1, 1, 0},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isOneBitCharacter(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
