package problem234

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	in   []int
	want bool
}

func TestIsPalindrome(t *testing.T) {
	tests := [...]testType{
		{
			in:   []int{1, 2},
			want: false,
		},
		{
			in:   []int{1, 2, 1},
			want: true,
		},
		{
			in:   []int{1, 2, 2, 1},
			want: true,
		},
	}
	for _, tt := range tests {
		got := isPalindrome(kit.SliceInt2ListNode(tt.in))
		if got != tt.want {
			t.Fatalf("in: %v, got: %v, want: %v", tt.in, got, tt.want)
		}
	}
}
