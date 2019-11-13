package kit

// ListNode - Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// ListNode2SliceInt - convert *ListNode to []int
func ListNode2SliceInt(l *ListNode) (s []int) {
	for l != nil {
		s = append(s, l.Val)
		l = l.Next
	}
	return
}

// SliceInt2ListNode - convert []int to *ListNode
func SliceInt2ListNode(s []int) *ListNode {
	l := &ListNode{}
	t := l
	for _, v := range s {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
