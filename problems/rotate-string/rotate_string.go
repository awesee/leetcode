package problem796

import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	A = strings.Repeat(A, 2)
	return strings.Contains(A, B)
}
