package problem606

import (
	"strconv"

	"github.com/openset/leetcode/internal/kit"
)

// TreeNode - Definition for a binary tree node.
type TreeNode = kit.TreeNode

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
