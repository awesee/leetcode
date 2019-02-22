package cousins_in_binary_tree

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	x        int
	y        int
	expected bool
}

func TestIsCousins(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{1, 2, 3, 4},
			x:        4,
			y:        3,
			expected: false,
		},
		{
			input:    []int{1, 2, 3, NULL, 4, NULL, 5},
			x:        5,
			y:        4,
			expected: true,
		},
		{
			input:    []int{1, 2, 3, NULL, 4},
			x:        2,
			y:        3,
			expected: false,
		},
		{
			input:    []int{1, 2, 3, NULL, 4, 5},
			x:        4,
			y:        5,
			expected: true,
		},
		{
			input:    []int{1, 2, 3, NULL, 4, 5},
			x:        5,
			y:        4,
			expected: true,
		},
	}
	for _, tc := range tests {
		output := isCousins(SliceInt2TreeNode(tc.input), tc.x, tc.y)
		if output != tc.expected {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
