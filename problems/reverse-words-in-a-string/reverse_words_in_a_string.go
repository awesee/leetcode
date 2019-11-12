package problem151

import (
	"regexp"
	"strings"
)

func reverseWords(s string) string {
	reg := regexp.MustCompile(`\S+`)
	words := reg.FindAllString(s, -1)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
