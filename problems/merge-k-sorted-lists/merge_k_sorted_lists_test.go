package problem23

import (
	"reflect"
	"testing"

	"github.com/awesee/leetcode/internal/kit"
)

type testType struct {
	in   [][]int
	want []int
}

func TestMergeKLists(t *testing.T) {
	tests := [...]testType{
		{
			in: [][]int{
				{1, 4, 5},
				{1, 3, 4},
				{2, 6},
			},
			want: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			in: [][]int{
				{2},
				{},
				{-1},
			},
			want: []int{-1, 2},
		},
		{
			in: [][]int{
				{},
				{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		lists := make([]*ListNode, len(tt.in))
		for i, v := range tt.in {
			lists[i] = kit.SliceInt2ListNode(v)
		}
		got := kit.ListNode2SliceInt(mergeKLists(lists))
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
