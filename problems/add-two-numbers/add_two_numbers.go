package add_two_numbers

import . "github.com/openset/leetcode/internal/kit"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	r := &ListNode{}
	t := r
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		t.Next = &ListNode{Val: val % 10}
		t = t.Next
		carry = val / 10
	}
	return r.Next
}
