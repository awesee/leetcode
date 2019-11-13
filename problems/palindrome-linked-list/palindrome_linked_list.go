package problem234

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
func isPalindrome(head *ListNode) bool {
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	for slow != nil {
		prev, slow, slow.Next = slow, slow.Next, prev
	}
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}
	return true
}
