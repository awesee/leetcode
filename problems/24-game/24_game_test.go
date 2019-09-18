package problem_679

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestJudgePoint24(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 1, 8, 7},
			expected: true,
		},
		{
			input:    []int{1, 2, 1, 2},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := judgePoint24(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
