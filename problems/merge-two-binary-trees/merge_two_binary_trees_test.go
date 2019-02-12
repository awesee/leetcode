package merge_two_binary_trees

import (
	"reflect"
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	t1       []int
	t2       []int
	expected []int
}

func TestMergeTrees(t *testing.T) {
	tests := [...]caseType{
		{
			t1:       []int{1, 3, 2, 5},
			t2:       []int{2, 1, 3, NULL, 4, NULL, 7},
			expected: []int{3, 4, 5, 5, 4, NULL, 7},
		},
	}
	for _, tc := range tests {
		t3 := mergeTrees(SliceInt2TreeNode(tc.t1), SliceInt2TreeNode(tc.t2))
		output := TreeNode2SliceInt(t3)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.t1, tc.t2, output, tc.expected)
		}
	}
}
