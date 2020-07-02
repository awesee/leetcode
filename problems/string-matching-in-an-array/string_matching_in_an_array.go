package problem1408

import "strings"

func stringMatching(words []string) []string {
	var ans []string
	for _, word := range words {
		for _, val := range words {
			if val != word && strings.Contains(val, word) {
				ans = append(ans, word)
				break
			}
		}
	}
	return ans
}
