package convert_sorted_array_to_binary_search_tree

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	input    []int
	expected []int
}

func TestSortedArrayToBST(t *testing.T) {
	tests := [...]caseType{
		{
			input:    []int{-10, -3, 0, 5, 9},
			expected: []int{0, -3, 9, -10, kit.NULL, 5},
		},
	}
	for _, tc := range tests {
		output := kit.TreeNode2SliceInt(sortedArrayToBST(tc.input))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
