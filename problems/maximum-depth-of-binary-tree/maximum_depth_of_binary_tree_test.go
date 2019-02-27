package maximum_depth_of_binary_tree

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected int
}

func TestMaxDepth(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 9, 20, NULL, NULL, 15, 7},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3, NULL, 5, 6},
			expected: 3,
		},
		{
			input:    []int{1, 2, NULL, NULL, 5},
			expected: 3,
		},
		{
			input:    []int{1, NULL, 3, NULL, 5},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := maxDepth(SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
