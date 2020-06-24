package problem1455

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	items := strings.Fields(sentence)
	for i, item := range items {
		if strings.HasPrefix(item, searchWord) {
			return i + 1
		}
	}
	return -1
}
