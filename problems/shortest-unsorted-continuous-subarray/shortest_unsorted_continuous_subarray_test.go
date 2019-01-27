package shortest_unsorted_continuous_subarray

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestFindUnsortedSubarray(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 6, 4, 8, 10, 9, 15},
			expected: 5,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			input:    []int{1, 2, 4, 5, 3},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := findUnsortedSubarray(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
