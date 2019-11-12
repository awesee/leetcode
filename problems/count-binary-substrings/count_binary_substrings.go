package problem696

func countBinarySubstrings(s string) int {
	ans, pre, cur, l := 0, 0, 1, len(s)
	for i := 1; i < l; i++ {
		if s[i] != s[i-1] {
			ans += min(pre, cur)
			pre = cur
			cur = 1
		} else {
			cur++
		}
	}
	return ans + min(pre, cur)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
