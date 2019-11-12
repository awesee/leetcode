package problem953

func isAlienSorted(words []string, order string) bool {
	m := make(map[byte]int)
	for i := len(order) - 1; i >= 0; i-- {
		m[order[i]] = i
	}
	if len(words) > 1 {
		for pre, word := range words[1:] {
			l, wl := len(words[pre]), len(word)
			for i := 0; i < l; i++ {
				if wl <= i || m[words[pre][i]] > m[word[i]] {
					return false
				} else if m[words[pre][i]] < m[word[i]] {
					break
				}
			}
		}
	}
	return true
}
