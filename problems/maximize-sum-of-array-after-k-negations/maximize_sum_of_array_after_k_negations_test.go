package maximize_sum_of_array_after_k_negations

import "testing"

type caseType struct {
	input    []int
	k        int
	expected int
}

func TestLargestSumAfterKNegations(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 2, 3},
			k:        1,
			expected: 5,
		},
		{
			input:    []int{3, -1, 0, 2},
			k:        3,
			expected: 6,
		},
		{
			input:    []int{2, -3, -1, 5, -4},
			k:        2,
			expected: 13,
		},
	}
	for _, tc := range tests {
		output := largestSumAfterKNegations(tc.input, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
