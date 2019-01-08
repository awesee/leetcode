package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	ws, pre, l := []byte(s+" "), 0, len(s)
	for cur, c := range ws {
		if c == ' ' {
			word := ws[pre:cur]
			for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
				word[i], word[j] = word[j], word[i]
			}
			pre = cur + 1
		}
	}
	return string(ws[:l])
}
