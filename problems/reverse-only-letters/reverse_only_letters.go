package problem917

import "unicode"

func reverseOnlyLetters(S string) string {
	s := []rune(S)
	i, j := 0, len(s)-1
	for i < j {
		if !unicode.IsLetter(rune(S[i])) {
			i++
		} else if !unicode.IsLetter(rune(S[j])) {
			j--
		} else {
			s[i], s[j] = s[j], s[i]
			i++
			j--
		}
	}
	return string(s)
}
