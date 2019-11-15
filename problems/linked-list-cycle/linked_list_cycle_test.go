package problem141

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	pos  int
	want bool
}

func TestHasCycle(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{3, 2, 0, -4},
			pos:  1,
			want: true,
		},
		{
			in:   []int{1, 2},
			pos:  0,
			want: true,
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			pos:  -1,
			want: false,
		},
		{
			in:   []int{1},
			pos:  -1,
			want: false,
		},
	}
	for _, tt := range tests {
		in := kit.SliceInt2ListNode(tt.in)
		p, curr := in, in
		for i := 0; curr != nil && tt.pos >= 0; i, curr = i+1, curr.Next {
			if i == tt.pos {
				p = curr
			}
			if curr.Next == nil {
				curr.Next = p
				break
			}
		}
		got := hasCycle(in)
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
