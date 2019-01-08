package reverse_words_in_a_string_iii

func reverseWords(s string) string {
	ss, pre, l := []byte(s+" "), 0, len(s)
	for cur, c := range ss {
		if c == ' ' {
			for i, j := pre, cur-1; i < j; i, j = i+1, j-1 {
				ss[i], ss[j] = ss[j], ss[i]
			}
			pre = cur + 1
		}
	}
	return string(ss[:l])
}
