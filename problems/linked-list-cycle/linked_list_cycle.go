package problem141

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
func hasCycle(head *ListNode) bool {
	for p1, p2 := head, head; p1 != nil && p2 != nil; p1 = p1.Next {
		if p2 = p2.Next; p2 != nil {
			if p2 = p2.Next; p1 == p2 {
				return true
			}
		}
	}
	return false
}
