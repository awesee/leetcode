package problem623

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	v    int
	d    int
	want []int
}

func TestAddOneRow(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 2, 6, 3, 1, 5},
			v:    1,
			d:    2,
			want: []int{4, 1, 1, 2, kit.NULL, kit.NULL, 6, 3, 1, 5},
		},
		{
			in:   []int{4, 2, kit.NULL, 3, 1},
			v:    1,
			d:    3,
			want: []int{4, 2, kit.NULL, 1, 1, 3, kit.NULL, kit.NULL, 1},
		},
		{
			in:   []int{3, 1},
			v:    1,
			d:    1,
			want: []int{1, 3, kit.NULL, 1},
		},
		{
			in:   []int{3, kit.NULL, 2},
			v:    1,
			d:    3,
			want: []int{3, kit.NULL, 2, 1, 1},
		},
	}
	for _, tt := range tests {
		root := kit.SliceInt2TreeNode(tt.in)
		root = addOneRow(root, tt.v, tt.d)
		got := kit.TreeNode2SliceInt(root)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
