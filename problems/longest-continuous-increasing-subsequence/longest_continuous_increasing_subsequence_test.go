package longest_continuous_increasing_subsequence

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestFindLengthOfLCIS(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 3, 5, 4, 7},
			expected: 3,
		},
		{
			input:    []int{2, 2, 2, 2, 2},
			expected: 1,
		},
		{
			input:    []int{},
			expected: 0,
		},
		{
			input:    []int{1, 3, 5, 7},
			expected: 4,
		},
	}
	for _, tc := range tests {
		output := findLengthOfLCIS(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
