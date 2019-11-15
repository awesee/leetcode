package problem993

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	x    int
	y    int
	want bool
}

func TestIsCousins(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4},
			x:    4,
			y:    3,
			want: false,
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 4, kit.NULL, 5},
			x:    5,
			y:    4,
			want: true,
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 4},
			x:    2,
			y:    3,
			want: false,
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 4, 5},
			x:    4,
			y:    5,
			want: true,
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 4, 5},
			x:    5,
			y:    4,
			want: true,
		},
	}
	for _, tt := range tests {
		got := isCousins(kit.SliceInt2TreeNode(tt.in), tt.x, tt.y)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
