package problem1408

import "strings"

func stringMatching(words []string) []string {
	var ans []string
	ok := func(str string) bool {
		for _, word := range words {
			if word != str && strings.Contains(word, str) {
				return true
			}
		}
		return false
	}
	for _, word := range words {
		if ok(word) {
			ans = append(ans, word)
		}
	}
	return ans
}
