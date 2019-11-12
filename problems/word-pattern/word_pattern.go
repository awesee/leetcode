package problem290

import "strings"

func wordPattern(pattern string, str string) bool {
	ss, l := strings.Split(str, " "), len(pattern)
	m, e := make(map[byte]string), make(map[string]bool)
	if len(ss) != l {
		return false
	}
	for i := 0; i < l; i++ {
		if s, ok := m[pattern[i]]; ok {
			if ss[i] != s {
				return false
			}
		} else if e[ss[i]] {
			return false
		} else {
			m[pattern[i]], e[ss[i]] = ss[i], true
		}
	}
	return true
}
