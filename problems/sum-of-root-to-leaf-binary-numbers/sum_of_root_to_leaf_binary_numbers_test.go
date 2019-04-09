package sum_of_root_to_leaf_binary_numbers

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected int
}

func TestSumRootToLeaf(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 0, 1, 0, 1, 0, 1},
			expected: 22,
		},
		{
			input:    []int{1, 1},
			expected: 3,
		},
		{
			input:    []int{1, NULL, 0},
			expected: 2,
		},
	}
	for _, tc := range tests {
		output := sumRootToLeaf(SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
