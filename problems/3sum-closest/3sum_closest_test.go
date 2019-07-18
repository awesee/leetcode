package p_3sum_closest

import "testing"

type caseType struct {
	input    []int
	target   int
	expected int
}

func TestThreeSumClosest(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := threeSumClosest(tc.input, tc.target)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
