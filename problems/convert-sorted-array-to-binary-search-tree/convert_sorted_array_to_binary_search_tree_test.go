package problem108

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want []int
}

func TestSortedArrayToBST(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{-10, -3, 0, 5, 9},
			want: []int{0, -3, 9, -10, kit.NULL, 5},
		},
	}
	for _, tt := range tests {
		got := kit.TreeNode2SliceInt(sortedArrayToBST(tt.in))
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
