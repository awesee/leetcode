package problem606

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want string
}

func TestTree2str(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4},
			want: "1(2(4))(3)",
		},
		{
			in:   []int{1, 2, 3, kit.NULL, 4},
			want: "1(2()(4))(3)",
		},
	}
	for _, tt := range tests {
		got := tree2str(kit.SliceInt2TreeNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
