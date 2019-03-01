package intersection_of_two_linked_lists

import (
	"testing"

	. "github.com/openset/leetcode/internal/kit"
)

type caseType struct {
	headA        []int
	headB        []int
	intersection []int
}

func TestGetIntersectionNode(t *testing.T) {
	tests := [...]caseType{
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
	for _, tc := range tests {
		intersection := SliceInt2ListNode(tc.intersection)
		headA := SliceInt2ListNode(tc.headA)
		headB := SliceInt2ListNode(tc.headB)
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
		output := getIntersectionNode(headA, headB)
		if output != intersection {
			t.Fatalf("input: %v, output: %v %v, expected: %v", tc.headA, tc.headB, output, tc.intersection)
		}
	}
}
