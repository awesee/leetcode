package merge_two_sorted_lists

import (
	"testing"
)

type caseType struct {
	l1       []int
	l2       []int
	expected []int
}

func TestMergeTwoLists(t *testing.T) {
	tests := [...]caseType{
		{
			l1:       []int{1, 2, 4},
			l2:       []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1:       nil,
			l2:       nil,
			expected: nil,
		},
		{
			l1:       nil,
			l2:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			l1:       []int{1, 2, 3},
			l2:       []int{1, 2, 3},
			expected: []int{1, 1, 2, 2, 3, 3},
		},
	}

	for _, tc := range tests {
		l1 := s2l(tc.l1)
		l2 := s2l(tc.l2)
		output := l2s(mergeTwoLists(l1, l2))
		if len(output) != len(tc.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", tc.l1, tc.l2, output, tc.expected)
		}
		for k, v := range tc.expected {
			if output[k] != v {
				t.Fatalf("input: %v %v, output: %v, expected: %v", tc.l1, tc.l2, output, tc.expected)
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
