package third_maximum_number

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestThirdMax(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 2, 5, 3, 5},
			expected: 2,
		},
		{
			input:    []int{1, 2, 2},
			expected: 2,
		},
		{
			input:    []int{1, 1, 2},
			expected: 2,
		},
		{
			input:    []int{3, 2, 1},
			expected: 1,
		},
		{
			input:    []int{1, 2},
			expected: 2,
		},
		{
			input:    []int{2, 2, 3, 1},
			expected: 1,
		},
	}
	for _, tc := range tests {
		output := thirdMax(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
