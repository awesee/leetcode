package merge_two_sorted_lists

import (
	"testing"
)

type caseType struct {
	L1       []int
	L2       []int
	Expected []int
}

func TestMergeTwoLists(t *testing.T) {
	tests := [...]caseType{
		{
			L1:       []int{1, 2, 4},
			L2:       []int{1, 3, 4},
			Expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			L1:       nil,
			L2:       nil,
			Expected: nil,
		},
		{
			L1:       nil,
			L2:       []int{1, 2, 3},
			Expected: []int{1, 2, 3},
		},
		{
			L1:       []int{1, 2, 3},
			L2:       []int{1, 2, 3},
			Expected: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, test := range tests {
		l1 := s2l(test.L1)
		l2 := s2l(test.L2)
		output := l2s(mergeTwoLists(l1, l2))
		if len(output) != len(test.Expected) {
			t.Fatalf("Input: %v %v, Output: %v, Expected: %v", test.L1, test.L2, output, test.Expected)
		}
		for k, v := range test.Expected {
			if output[k] != v {
				t.Fatalf("Input: %v %v, Output: %v, Expected: %v", test.L1, test.L2, output, test.Expected)
			}
		}
	}
}

// convert []int to *ListNode
func s2l(s []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range s {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// convert *ListNode to []int
func l2s(l *ListNode) (s []int) {
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return
}
