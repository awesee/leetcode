package problem215

import "testing"

type caseType struct {
	input    []int
	k        int
	expected int
}

func TestFindKthLargest(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		}, {
			input:    []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
	}
	for _, tc := range tests {
		output := findKthLargest(tc.input, tc.k)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
