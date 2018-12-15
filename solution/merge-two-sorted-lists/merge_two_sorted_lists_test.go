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

	for _, test := range tests {
		l1 := s2l(test.l1)
		l2 := s2l(test.l2)
		output := l2s(mergeTwoLists(l1, l2))
		if len(output) != len(test.expected) {
			t.Fatalf("input: %v %v, output: %v, expected: %v", test.l1, test.l2, output, test.expected)
		}
		for k, v := range test.expected {
			if output[k] != v {
				t.Fatalf("input: %v %v, output: %v, expected: %v", test.l1, test.l2, output, test.expected)
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
