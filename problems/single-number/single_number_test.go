package single_number

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestSingleNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 2, 1},
			expected: 1,
		},
		{
			input:    []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			input:    []int{1, 2, 1, 2, 3},
			expected: 3,
		},
	}

	for _, tc := range tests {
		output := singleNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
