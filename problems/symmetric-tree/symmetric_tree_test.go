package problem101

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected bool
}

func TestIsSymmetric(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 2, 3, 4, 4, 3},
			expected: true,
		},
		{
			input:    []int{1, 2, 2, kit.NULL, 3, kit.NULL, 3},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isSymmetric(kit.SliceInt2TreeNode(tc.input))
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
