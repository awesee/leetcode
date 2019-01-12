package degree_of_an_array

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestFindShortestSubArray(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 2, 3, 1},
			expected: 2,
		},
		{
			input:    []int{1, 2, 2, 3, 1, 4, 2},
			expected: 6,
		},
		{
			input:    []int{1},
			expected: 1,
		},
		{
			input:    []int{2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2},
			expected: 7,
		},
		{
			input:    []int{1, 2, 2, 3, 1},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := findShortestSubArray(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
