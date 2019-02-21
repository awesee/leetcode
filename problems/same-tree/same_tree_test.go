package same_tree

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	p        []int
	q        []int
	expected bool
}

func TestIsSameTree(t *testing.T) {
	tests := [...]caseType{
		{
			p:        []int{1, 2, 3},
			q:        []int{1, 2, 3},
			expected: true,
		},
		{
			p:        []int{1, 2},
			q:        []int{1, NULL, 2},
			expected: false,
		},
		{
			p:        []int{1, 2, 1},
			q:        []int{1, 1, 2},
			expected: false,
		},
	}
	for _, tc := range tests {
		output := isSameTree(SliceInt2TreeNode(tc.p), SliceInt2TreeNode(tc.q))
		if output != tc.expected {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.p, tc.q, output, tc.expected)
		}
	}
}
