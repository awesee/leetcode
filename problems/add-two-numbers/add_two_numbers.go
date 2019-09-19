package problem_2

import "github.com/openset/leetcode/internal/kit"

type ListNode = kit.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans, carry := &ListNode{}, 0
	t := ans
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
		carry = val / 10
		t.Next = &ListNode{Val: val % 10}
		t = t.Next
	}
	return ans.Next
}
