package problem83

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
func deleteDuplicates(head *ListNode) *ListNode {
	for prev, t := head, head; t != nil; t = t.Next {
		if prev.Val == t.Val {
			prev.Next = t.Next
		} else {
			prev = t
		}
	}
	return head
}
