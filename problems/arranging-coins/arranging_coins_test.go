package problem_441

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestArrangeCoins(t *testing.T) {
	tests := [...]caseType{
		{
			input:    5,
			expected: 2,
		},
		{
			input:    8,
			expected: 3,
		},
		{
			input:    0,
			expected: 0,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 1,
		},
		{
			input:    3,
			expected: 2,
		},
		{
			input:    13,
			expected: 4,
		},
		{
			input:    130,
			expected: 15,
		},
	}
	for _, tc := range tests {
		output := arrangeCoins(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
