package range_sum_of_bst

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	l        int
	r        int
	expected int
}

func TestRangeSumBST(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{10, 5, 15, 3, 7, NULL, 18},
			l:        7,
			r:        15,
			expected: 32,
		},
		{
			input:    []int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6},
			l:        6,
			r:        10,
			expected: 23,
		},
	}
	for _, tc := range tests {
		output := rangeSumBST(SliceInt2TreeNode(tc.input), tc.l, tc.r)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
