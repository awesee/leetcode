package add_one_row_to_tree

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	v        int
	d        int
	expected []int
}

func TestAddOneRow(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{4, 2, 6, 3, 1, 5},
			v:        1,
			d:        2,
			expected: []int{4, 1, 1, 2, kit.NULL, kit.NULL, 6, 3, 1, 5},
		},
		{
			input:    []int{4, 2, kit.NULL, 3, 1},
			v:        1,
			d:        3,
			expected: []int{4, 2, kit.NULL, 1, 1, 3, kit.NULL, kit.NULL, 1},
		},
		{
			input:    []int{3, 1},
			v:        1,
			d:        1,
			expected: []int{1, 3, kit.NULL, 1},
		},
		{
			input:    []int{3, kit.NULL, 2},
			v:        1,
			d:        3,
			expected: []int{3, kit.NULL, 2, 1, 1},
		},
	}
	for _, tc := range tests {
		root := kit.SliceInt2TreeNode(tc.input)
		root = addOneRow(root, tc.v, tc.d)
		output := kit.TreeNode2SliceInt(root)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
