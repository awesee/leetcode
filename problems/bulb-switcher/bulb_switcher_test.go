package problem319

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestBulbSwitch(t *testing.T) {
	tests := [...]caseType{
		{
			input:    3,
			expected: 1,
		},
		{
			input:    4,
			expected: 2,
		},
		{
			input:    5,
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := bulbSwitch(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
