package monotonic_array

import "testing"

type caseType struct {
	input    []int
	expected bool
}

func TestIsMonotonic(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 2, 3},
			expected: true,
		},
		{
			input:    []int{6, 5, 4, 4},
			expected: true,
		},
		{
			input:    []int{1, 3, 2},
			expected: false,
		},
		{
			input:    []int{1, 2, 4, 5},
			expected: true,
		},
		{
			input:    []int{1, 1, 1},
			expected: true,
		},
		{
			input:    []int{1, 2, 3, 3, 3, 2, 1},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isMonotonic(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
