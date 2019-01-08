package repeated_substring_pattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	return strings.Contains((s + s)[1:2*len(s)-1], s)
}
