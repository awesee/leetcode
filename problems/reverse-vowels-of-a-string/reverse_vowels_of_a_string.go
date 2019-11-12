package problem345

func reverseVowels(s string) string {
	i, j, ss := 0, len(s)-1, []byte(s)
	for i < j {
		if !isVowel(ss[i]) {
			i++
		} else if !isVowel(ss[j]) {
			j--
		} else {
			ss[i], ss[j] = ss[j], ss[i]
			i++
			j--
		}
	}
	return string(ss)
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}
