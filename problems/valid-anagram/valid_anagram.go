package problem242

func isAnagram(s string, t string) bool {
	sl := len(s)
	tl := len(t)
	if sl != tl {
		return false
	}
	var sc [256]int
	for i := 0; i < sl; i++ {
		sc[s[i]]++
	}
	for i := 0; i < tl; i++ {
		sc[t[i]]--
		if sc[t[i]] < 0 {
			return false
		}
	}
	return true
}
