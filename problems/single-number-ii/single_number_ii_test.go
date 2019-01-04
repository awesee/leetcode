package single_number_ii

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestSingleNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 2, 3, 2},
			expected: 3,
		},
		{
			input:    []int{0, 1, 0, 1, 0, 1, 99},
			expected: 99,
		},
	}
	for _, tc := range tests {
		output := singleNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
