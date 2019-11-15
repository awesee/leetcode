package problem160

import (
	"testing"

	"github.com/openset/leetcode/internal/kit"
)

type testType struct {
	headA        []int
	headB        []int
	intersection []int
}

func TestGetIntersectionNode(t *testing.T) {
	tests := [...]testType{
		{
			headA:        []int{4, 1},
			headB:        []int{5, 0, 1},
			intersection: []int{8, 4, 5},
		},
		{
			headA:        []int{0, 9, 1, 2},
			headB:        []int{3},
			intersection: []int{2, 4},
		},
		{
			headA:        []int{2, 4, 6},
			headB:        []int{1, 5},
			intersection: []int{},
		},
		{
			headA:        []int{1},
			headB:        []int{},
			intersection: []int{},
		},
	}
	for _, tt := range tests {
		intersection := kit.SliceInt2ListNode(tt.intersection)
		headA := kit.SliceInt2ListNode(tt.headA)
		headB := kit.SliceInt2ListNode(tt.headB)
		for n := headA; n != nil; n = n.Next {
			if n.Next == nil {
				n.Next = intersection
				break
			}
		}
		for n := headB; n != nil; n = n.Next {
			if n.Next == nil {
				n.Next = intersection
				break
			}
		}
		got := getIntersectionNode(headA, headB)
		if got != intersection {
			t.Fatalf("in: %v, got: %v %v, want: %v", tt.headA, tt.headB, got, tt.intersection)
		}
	}
}
