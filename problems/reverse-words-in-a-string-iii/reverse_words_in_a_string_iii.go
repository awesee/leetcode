package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	ss, pre, l := []byte(s+" "), 0, len(s)
	for cur, c := range ss {
		if c == ' ' {
			ws := ss[pre:cur]
			for i, j := 0, len(ws)-1; i < j; i, j = i+1, j-1 {
				ws[i], ws[j] = ws[j], ws[i]
			}
			pre = cur + 1
		}
	}
	return string(ss[:l])
}
