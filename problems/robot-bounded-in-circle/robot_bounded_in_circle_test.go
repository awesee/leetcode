package robot_bounded_in_circle

import "testing"

type caseType struct {
	input    string
	expected bool
}

func TestIsRobotBounded(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "GGLLGG",
			expected: true,
		},
		{
			input:    "GG",
			expected: false,
		},
		{
			input:    "GL",
			expected: true,
		},
		{
			input:    "GGLLGGGGRRGG",
			expected: true,
		},
		{
			input:    "GGRRGG",
			expected: true,
		},
		{
			input:    "GLRLLGLL",
			expected: true,
		},
		{
			input:    "GLGRGLGLGLGL",
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isRobotBounded(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
