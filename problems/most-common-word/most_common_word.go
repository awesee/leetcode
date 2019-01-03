package most_common_word

import "strings"

func mostCommonWord(paragraph string, banned []string) string {
	word, words := "", make([]string, 0)
	paragraph = strings.ToLower(paragraph) + " "
	for _, char := range paragraph {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			word += string(char)
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	m := make(map[string]bool)
	for _, ban := range banned {
		m[ban] = true
	}
	ans, max, count := "", 0, make(map[string]int)
	for _, word = range words {
		if !m[word] {
			count[word]++
			if count[word] > max {
				max = count[word]
				ans = word
			}
		}
	}
	return ans
}
