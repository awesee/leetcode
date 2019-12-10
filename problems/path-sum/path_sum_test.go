package problem112

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	sum  int
	want bool
}

func TestHasPathSum(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{5, 4, 8, 11, kit.NULL, 13, 4, 7, 2, kit.NULL, kit.NULL, kit.NULL, 1},
			sum:  22,
			want: true,
		},
	}
	for _, tt := range tests {
		got := hasPathSum(kit.SliceInt2TreeNode(tt.in), tt.sum)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
