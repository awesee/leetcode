package problem206

import (
	"reflect"
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want []int
}

func TestReverseList(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2, 3, 4, 5},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		got := kit.ListNode2SliceInt(reverseList(kit.SliceInt2ListNode(tt.in)))
		if !reflect.DeepEqual(got, tt.want) {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
