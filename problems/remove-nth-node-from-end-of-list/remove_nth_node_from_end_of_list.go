package problem19

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	q := []*ListNode{head}
	for head.Next != nil {
		q = append(q, head.Next)
		head = head.Next
	}
	i := len(q) - n
	if i > 0 {
		q[i-1].Next = q[i].Next
	} else {
		q[0] = q[i].Next
	}
	return q[0]
}
