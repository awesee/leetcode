package maximum_product_of_three_numbers

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestMaximumProduct(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3},
			expected: 6,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: 24,
		},
		{
			input:    []int{-2, -1, 0, 2, 3},
			expected: 6,
		},
	}
	for _, tc := range tests {
		output := maximumProduct(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
