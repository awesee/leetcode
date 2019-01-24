package missing_number

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMissingNumber(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			input:    []int{0, 1, 2, 4, 5, 6},
			expected: 3,
		},
		{
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}
	for _, tc := range tests {
		output := missingNumber(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
