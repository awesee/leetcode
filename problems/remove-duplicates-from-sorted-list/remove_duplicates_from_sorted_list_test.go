package problem83

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want []int
}

func TestDeleteDuplicates(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 1, 2},
			want: []int{1, 2},
		},
		{
			in:   []int{1, 1, 2, 3, 3},
			want: []int{1, 2, 3},
		},
		{
			in:   []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 1, 2},
		},
	}
	for _, tt := range tests {
		got := deleteDuplicates(kit.SliceInt2ListNode(tt.in))
		if !reflect.DeepEqual(kit.ListNode2SliceInt(got), tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
