package problem824

import "strings"

func toGoatLatin(S string) string {
	ans := ""
	s := strings.Split(S, " ")
	for i, word := range s {
		switch word[0] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ans += " " + word + "ma"
		default:
			ans += " " + word[1:] + word[0:1] + "ma"
		}
		ans += strings.Repeat("a", i+1)
	}
	return ans[1:]
}
