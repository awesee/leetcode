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
	ans := &ListNode{Next: head}
	pos := ans
	for head != nil {
		head = head.Next
		if n > 0 {
			n--
		} else {
			pos = pos.Next
		}
	}
	pos.Next = pos.Next.Next
	return ans.Next
}
