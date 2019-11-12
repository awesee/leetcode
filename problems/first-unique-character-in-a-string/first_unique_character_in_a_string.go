package problem387

func firstUniqChar(s string) int {
	var count [256]byte
	l := len(s)
	for i := 0; i < l; i++ {
		if count[s[i]] < 2 {
			count[s[i]]++
		}
	}
	for i := 0; i < l; i++ {
		if count[s[i]] == 1 {
			return i
		}
	}
	return -1
}
