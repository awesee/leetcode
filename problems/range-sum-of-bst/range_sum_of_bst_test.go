package problem938

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	l    int
	r    int
	want int
}

func TestRangeSumBST(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{10, 5, 15, 3, 7, kit.NULL, 18},
			l:    7,
			r:    15,
			want: 32,
		},
		{
			in:   []int{10, 5, 15, 3, 7, 13, 18, 1, kit.NULL, 6},
			l:    6,
			r:    10,
			want: 23,
		},
	}
	for _, tt := range tests {
		got := rangeSumBST(kit.SliceInt2TreeNode(tt.in), tt.l, tt.r)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
