package problem101

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want bool
}

func TestIsSymmetric(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 2, 3, 4, 4, 3},
			want: true,
		},
		{
			in:   []int{1, 2, 2, kit.NULL, 3, kit.NULL, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isSymmetric(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
