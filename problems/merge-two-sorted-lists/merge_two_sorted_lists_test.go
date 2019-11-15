package problem21

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	l1   []int
	l2   []int
	want []int
}

func TestMergeTwoLists(t *testing.T) {
	tests := [...]testType{
		{
			l1:   []int{1, 2, 4},
			l2:   []int{1, 3, 4},
			want: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:   nil,
			l2:   nil,
			want: nil,
		},
		{
			l1:   nil,
			l2:   []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			l1:   []int{1, 2, 3},
			l2:   nil,
			want: []int{1, 2, 3},
		},
		{
			l1:   []int{1, 2, 3, 8, 16},
			l2:   []int{1, 2, 3, 10, 12, 24},
			want: []int{1, 1, 2, 2, 3, 3, 8, 10, 12, 16, 24},
		},
	}
	for _, tt := range tests {
		l1 := kit.SliceInt2ListNode(tt.l1)
		l2 := kit.SliceInt2ListNode(tt.l2)
		got := kit.ListNode2SliceInt(mergeTwoLists(l1, l2))
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v %v, got: %v, want: %v", tt.l1, tt.l2, got, tt.want)
		}
	}
}
