package problem1662

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := strings.Builder{}
	for _, s := range word1 {
		str1.WriteString(s)
	}
	str2 := strings.Builder{}
	for _, s := range word2 {
		str2.WriteString(s)
	}
	return str1.String() == str2.String()
}
