package array_partition_i

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestArrayPairSum(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 4, 3, 2},
			expected: 4,
		},
	}
	for _, tc := range tests {
		output := arrayPairSum(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
