package problem104

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want int
}

func TestMaxDepth(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 9, 20, kit.NULL, kit.NULL, 15, 7},
			want: 3,
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 5, 6},
			want: 3,
		},
		{
			in:   []int{1, 2, kit.NULL, kit.NULL, 5},
			want: 3,
		},
		{
			in:   []int{1, kit.NULL, 3, kit.NULL, 5},
			want: 3,
		},
	}
	for _, tt := range tests {
		got := maxDepth(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
