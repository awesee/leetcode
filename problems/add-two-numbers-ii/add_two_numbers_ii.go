package problem445

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	ans, n1, n2, carry := &ListNode{}, len(s1)-1, len(s2)-1, 0
	for n1 >= 0 || n2 >= 0 || carry > 0 {
		if n1 >= 0 {
			carry += s1[n1]
			n1--
		}
		if n2 >= 0 {
			carry += s2[n2]
			n2--
		}
		ans.Val = carry % 10
		ans = &ListNode{Val: 0, Next: ans}
		carry /= 10
	}
	return ans.Next
}
