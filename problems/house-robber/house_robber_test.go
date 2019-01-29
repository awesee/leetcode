package house_robber

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestRob(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
	}
	for _, tc := range tests {
		output := rob(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
