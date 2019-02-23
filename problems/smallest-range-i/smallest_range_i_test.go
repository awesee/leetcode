package smallest_range_i

import "testing"

type caseType struct {
	input    []int
	k        int
	expected int
}

func TestSmallestRangeI(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1},
			k:        0,
			expected: 0,
		},
		{
			input:    []int{0, 10},
			k:        2,
			expected: 6,
		},
		{
			input:    []int{1, 3, 6},
			k:        3,
			expected: 0,
		},
		{
			input:    []int{1, 3, 5, 7, 9},
			k:        1,
			expected: 6,
		},
		{
			input:    []int{18, 16, 12, 7, 9, 3, 5},
			k:        6,
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := smallestRangeI(tc.input, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
