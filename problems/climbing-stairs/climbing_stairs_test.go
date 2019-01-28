package climbing_stairs

import "testing"

type caseType struct {
	input    int
	expected int
}

func TestClimbStairs(t *testing.T) {
	tests := [...]caseType{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := climbStairs(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
