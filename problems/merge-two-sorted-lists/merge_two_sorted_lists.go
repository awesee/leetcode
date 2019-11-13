package problem21

import "github.com/openset/leetcode/internal/kit"

// ListNode - Definition for singly-linked list.
type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}
