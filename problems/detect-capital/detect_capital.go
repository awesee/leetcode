package problem520

import "unicode"

func detectCapitalUse(word string) bool {
	count := 0
	for _, c := range word {
		if unicode.IsUpper(c) {
			count++
		}
	}
	return count == 0 || count == len(word) || (count == 1 && unicode.IsUpper(rune(word[0])))
}
