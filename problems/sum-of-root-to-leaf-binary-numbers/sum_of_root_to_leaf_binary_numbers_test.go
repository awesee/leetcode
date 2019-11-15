package problem1022

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want int
}

func TestSumRootToLeaf(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 0, 1, 0, 1, 0, 1},
			want: 22,
		},
		{
			in:   []int{1, 1},
			want: 3,
		},
		{
			in:   []int{1, kit.NULL, 0},
			want: 2,
		},
	}
	for _, tt := range tests {
		got := sumRootToLeaf(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
