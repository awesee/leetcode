package problem38

func countAndSay(n int) string {
	r := "1"
	for i := 1; i < n; i++ {
		r = say(r)
	}
	return r
}

func say(s string) string {
	r := ""
	var sv rune
	var sc byte
	for i, v := range s {
		if i == 0 {
			sv = v
			sc = '1'
		} else {
			if v == rune(s[i-1]) {
				sc++
			} else {
				r += string(sc) + string(sv)
				sv = v
				sc = '1'
			}
		}
	}
	if sc > '0' {
		r += string(sc) + string(sv)
	}
	return r
}
