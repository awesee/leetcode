package problem237

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	node int
	want []int
}

func TestDeleteNode(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{4, 5, 1, 9},
			node: 5,
			want: []int{4, 1, 9},
		},
		{
			in:   []int{4, 5, 1, 9},
			node: 1,
			want: []int{4, 5, 9},
		},
		{
			in:   []int{4, 5, 1, 9},
			node: 4,
			want: []int{5, 1, 9},
		},
		{
			in:   []int{4, 5, 1, 9},
			node: 9,
			want: []int{4, 5, 1, 9},
		},
		{
			in:   []int{4, 5, 1, 9},
			node: 2,
			want: []int{4, 5, 1, 9},
		},
	}
	for _, tt := range tests {
		head := kit.SliceInt2ListNode(tt.in)
		node := head
		for node != nil && node.Val != tt.node {
			node = node.Next
		}
		deleteNode(node)
		got := kit.ListNode2SliceInt(head)
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
