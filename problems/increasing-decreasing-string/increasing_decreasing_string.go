package problem1370

func sortString(s string) string {
	var m [26]int
	for _, c := range s {
		m[c-'a']++
	}
	l := len(s)
	ans := make([]byte, 0, l)
	appendChar := func(c int) {
		if m[c] > 0 {
			m[c]--
			l--
			ans = append(ans, byte(c+'a'))
		}
	}
	for l > 0 {
		for i := 0; i <= 25; i++ {
			appendChar(i)
		}
		for i := 25; i >= 0; i-- {
			appendChar(i)
		}
	}
	return string(ans)
}
