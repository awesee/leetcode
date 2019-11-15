package problem617

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	t1   []int
	t2   []int
	want []int
}

func TestMergeTrees(t *testing.T) {
	tests := [...]testType{
		{
			t1:   []int{1, 3, 2, 5},
			t2:   []int{2, 1, 3, kit.NULL, 4, kit.NULL, 7},
			want: []int{3, 4, 5, 5, 4, kit.NULL, 7},
		},
	}
	for _, tt := range tests {
		t3 := mergeTrees(kit.SliceInt2TreeNode(tt.t1), kit.SliceInt2TreeNode(tt.t2))
		got := kit.TreeNode2SliceInt(t3)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.t1, tt.t2, got, tt.want)
		}
	}
}
