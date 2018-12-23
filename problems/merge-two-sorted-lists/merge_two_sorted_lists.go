package merge_two_sorted_lists

import . "github.com/openset/leetcode/problems/000000"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{}
	t := l
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}
	if l1 == nil {
		t.Next = l2
	} else {
		t.Next = l1
	}
	return l.Next
}
