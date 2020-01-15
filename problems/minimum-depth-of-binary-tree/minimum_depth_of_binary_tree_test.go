package problem111

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want int
}

func TestMinDepth(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			want: 2,
		},
		{
			in:   []int{1, 2},
			want: 2,
		},
		{
			in:   nil,
			want: 0,
		},
	}
	for _, tt := range tests {
		got := minDepth(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
