package problem104

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected int
}

func TestMaxDepth(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			expected: 3,
		},
		{
			input:    []int{1, 2, 3, kit.NULL, 5, 6},
			expected: 3,
		},
		{
			input:    []int{1, 2, kit.NULL, kit.NULL, 5},
			expected: 3,
		},
		{
			input:    []int{1, kit.NULL, 3, kit.NULL, 5},
			expected: 3,
		},
	}
	for _, tc := range tests {
		output := maxDepth(kit.SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
