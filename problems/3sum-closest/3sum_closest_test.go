package problem_16

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
		{
			input:    []int{-1, 0, 1, 2, -1, -4},
			target:   1,
			expected: 1,
		},
		{
			input:    []int{0, 0, 0, 0},
			target:   1,
			expected: 0,
		},
		{
			input:    []int{-2, 0, 0, 2, 2, 2},
			target:   2,
			expected: 2,
		},
		{
			input:    []int{-2, 0, 0, 2, 2, 2, 2},
			target:   1,
			expected: 0,
		},
	}
	for _, tc := range tests {
		output := threeSumClosest(tc.input, tc.target)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
