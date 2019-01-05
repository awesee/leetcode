package kit

import "testing"

func TestTreeNode(t *testing.T) {
	tests := [...][]int{
		{},
	}
	for _, input := range tests {
		TreeNode2SliceInt(nil)
		SliceInt2TreeNode(input)
	}
}
