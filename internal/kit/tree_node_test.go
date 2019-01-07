package kit

import (
	"reflect"
	"testing"
)

func TestTreeNode(t *testing.T) {
	tests := [...]struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{1, 2, 3, NULL, 5},
			expected: []int{1, 2, 3, NULL, 5},
		},
		{
			input:    []int{1, 2, 3, NULL, 5, NULL, NULL, NULL},
			expected: []int{1, 2, 3, NULL, 5},
		},
		{
			input:    nil,
			expected: nil,
		},
	}
	for _, tc := range tests {
		output := TreeNode2SliceInt(SliceInt2TreeNode(tc.input))
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fatalf("input: %v, output: %v, expected: %v", tc.input, output, tc.expected)
		}
	}
}
