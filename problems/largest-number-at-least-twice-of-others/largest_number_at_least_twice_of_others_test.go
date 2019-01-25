package largest_number_at_least_twice_of_others

import "testing"

type caseType struct {
	input    []int
	expected int
}

func TestDominantIndex(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 6, 1, 0},
			expected: 1,
		},
		{
			input:    []int{1, 2, 3, 4},
			expected: -1,
		},
		{
			input:    []int{1},
			expected: 0,
		},
		{
			input:    []int{0, 0, 3, 2},
			expected: -1,
		},
		{
			input:    []int{3, 0, 0, 2},
			expected: -1,
		},
	}
	for _, tc := range tests {
		output := dominantIndex(tc.input)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
