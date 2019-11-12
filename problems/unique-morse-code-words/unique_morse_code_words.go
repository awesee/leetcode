package problem804

func uniqueMorseRepresentations(words []string) int {
	list := [...]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	m := make(map[string]bool)
	for _, word := range words {
		k := ""
		for i := 0; i < len(word); i++ {
			k += list[word[i]-'a']
		}
		m[k] = true
	}
	return len(m)
}
