package problem226

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want []int
}

func TestInvertTree(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 2, 7, 1, 3, 6, 9},
			want: []int{4, 7, 2, 9, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		out := invertTree(kit.SliceInt2TreeNode(tt.in))
		got := kit.TreeNode2SliceInt(out)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
