package maximum_subarray

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMaxSubArray(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
	}
	for _, tc := range tests {
		output := maxSubArray(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
