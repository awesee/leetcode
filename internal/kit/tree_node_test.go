package kit

import "testing"

func TestTreeNode(t *testing.T) {
	tests := [...][]int{
		{1, 2, 3, 4},
	}
	for _, input := range tests {
		TreeNode2SliceInt(nil)
		SliceInt2TreeNode(input)
	}
}
