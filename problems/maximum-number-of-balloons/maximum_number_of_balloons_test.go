package problem1189

import "testing"

type caseType struct {
	input    string
	expected int
}

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := [...]caseType{
		{
			input:    "nlaebolko",
			expected: 1,
		},
		{
			input:    "loonbalxballpoon",
			expected: 2,
		},
		{
			input:    "leetcode",
			expected: 0,
		},
		{
			input:    "lloo",
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := maxNumberOfBalloons(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
