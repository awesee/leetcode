package problem19

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	n    int
	want []int
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4, 5},
			n:    2,
			want: []int{1, 2, 3, 5},
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			n:    5,
			want: []int{2, 3, 4, 5},
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			n:    1,
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		got := removeNthFromEnd(kit.SliceInt2ListNode(tt.in), tt.n)
		if !reflect.DeepEqual(kit.ListNode2SliceInt(got), tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
