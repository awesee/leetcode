package univalued_binary_tree

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected bool
}

func TestIsUnivalTree(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 1, 1, 1, 1, NULL, 1},
			expected: true,
		},
		{
			input:    []int{2, 2, 2, 5, 2},
			expected: false,
		},
		{
			input:    []int{3, 3, 3, 3, 5},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isUnivalTree(SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
