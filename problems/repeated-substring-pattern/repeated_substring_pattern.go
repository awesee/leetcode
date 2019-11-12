package problem459

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:2*len(s)-1], s)
}
