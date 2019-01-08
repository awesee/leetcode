package repeated_substring_pattern

import "strings"

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	for i := 1; i <= l/2; i++ {
		if l%i == 0 && strings.Count(s, s[:i]) == l/i {
			return true
		}
	}
	return false
}
