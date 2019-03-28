package partition_array_into_three_parts_with_equal_sum

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestCanThreePartsEqualSum(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			expected: true,
		},
		{
			input:    []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			expected: false,
		},
		{
			input:    []int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4},
			expected: true,
		},
		{
			input:    []int{-3, 3, 3, -3},
			expected: false,
		},
		{
			input:    []int{1, 7, 2, 6, 3, 5, 8},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := canThreePartsEqualSum(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
