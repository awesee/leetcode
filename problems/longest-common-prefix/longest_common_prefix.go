package problem14

import "strings"

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < l; i++ {
		for strings.Index(strs[i], pre) != 0 {
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
