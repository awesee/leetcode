package problem_650

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestMinSteps(t *testing.T) {
	tests := [...]caseType{
		{
			input:    1,
			expected: 0,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    30,
			expected: 10,
		},
		{
			input:    97,
			expected: 97,
		},
	}
	for _, tc := range tests {
		output := minSteps(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
