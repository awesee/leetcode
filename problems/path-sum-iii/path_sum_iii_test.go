package problem437

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	sum  int
	want int
}

func TestPathSum(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{10, 5, -3, 3, 2, kit.NULL, 11, 3, -2, kit.NULL, 1},
			sum:  8,
			want: 3,
		},
	}
	for _, tt := range tests {
		root := kit.SliceInt2TreeNode(tt.in)
		got := pathSum(root, tt.sum)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
