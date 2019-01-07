package construct_string_from_binary_tree

import (
	"strconv"

	. "github.com/openset/leetcode/internal/kit"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	ans := ""
	if t != nil {
		ans += strconv.Itoa(t.Val)
		if t.Left != nil {
			ans += "(" + tree2str(t.Left) + ")"
		} else if t.Right != nil {
			ans += "()"
		}
		if t.Right != nil {
			ans += "(" + tree2str(t.Right) + ")"
		}
	}
	return ans
}
