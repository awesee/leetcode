package implement_strstr

import "strings"

func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
