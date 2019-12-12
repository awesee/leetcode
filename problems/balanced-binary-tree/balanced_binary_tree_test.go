package problem110

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want bool
}

func TestIsBalanced(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			want: true,
		},
		{
			in:   []int{1, 2, 2, 3, 3, kit.NULL, kit.NULL, 4, 4},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isBalanced(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
