package problem819

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	words := make([]string, 0)
	paragraph = strings.ToLower(paragraph)
	reg := regexp.MustCompile(`[a-z]+`)
	for _, word := range reg.FindAllString(paragraph, -1) {
		words = append(words, word)
	}
	mb := make(map[string]bool)
	for _, ban := range banned {
		mb[ban] = true
	}
	ans, max, count := "", 0, make(map[string]int)
	for _, word := range words {
		if !mb[word] {
			count[word]++
			if count[word] > max {
				max = count[word]
				ans = word
			}
		}
	}
	return ans
}
