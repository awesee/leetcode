package problem100

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	p    []int
	q    []int
	want bool
}

func TestIsSameTree(t *testing.T) {
	tests := [...]testType{
		{
			p:    []int{1, 2, 3},
			q:    []int{1, 2, 3},
			want: true,
		},
		{
			p:    []int{1, 2},
			q:    []int{1, kit.NULL, 2},
			want: false,
		},
		{
			p:    []int{1, 2, 1},
			q:    []int{1, 1, 2},
			want: false,
		},
	}
	for _, tt := range tests {
		got := isSameTree(kit.SliceInt2TreeNode(tt.p), kit.SliceInt2TreeNode(tt.q))
		if got != tt.want {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.p, tt.q, got, tt.want)
		}
	}
}
