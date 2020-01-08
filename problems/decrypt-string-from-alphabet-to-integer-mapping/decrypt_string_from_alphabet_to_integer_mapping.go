package problem1309

func freqAlphabets(s string) string {
	l := len(s)
	ans, head := make([]byte, l), l-1
	for i := l - 1; i >= 0; i-- {
		if s[i] == '#' {
			ans[head] = 'a' + (s[i-1] - '1') + (s[i-2]-'0')*10
			i -= 2
		} else {
			ans[head] = 'a' + (s[i] - '1')
		}
		head--
	}
	return string(ans[head+1:])
}
