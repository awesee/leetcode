package largest_perimeter_triangle

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestLargestPerimeter(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{2, 1, 2},
			expected: 5,
		},
		{
			input:    []int{1, 2, 1},
			expected: 0,
		},
		{
			input:    []int{3, 2, 3, 4},
			expected: 10,
		},
		{
			input:    []int{3, 6, 2, 3},
			expected: 8,
		},
	}
	for _, tc := range tests {
		output := largestPerimeter(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
