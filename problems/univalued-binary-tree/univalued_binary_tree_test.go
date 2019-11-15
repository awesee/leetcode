package problem965

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want bool
}

func TestIsUnivalTree(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 1, 1, 1, 1, kit.NULL, 1},
			want: true,
		},
		{
			in:   []int{2, 2, 2, 5, 2},
			want: false,
		},
		{
			in:   []int{3, 3, 3, 3, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isUnivalTree(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
