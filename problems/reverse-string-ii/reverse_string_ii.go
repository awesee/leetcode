package problem541

func reverseStr(s string, k int) string {
	ss, l := []rune(s), len(s)-1
	for i := 0; i < l; i += 2 * k {
		for m, n := i, i+k-1; m < n; m, n = m+1, n-1 {
			if n > l {
				n = l
			}
			ss[m], ss[n] = ss[n], ss[m]
		}
	}
	return string(ss)
}
