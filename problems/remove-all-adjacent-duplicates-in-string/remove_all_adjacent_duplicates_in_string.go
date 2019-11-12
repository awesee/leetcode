package problem1047

func removeDuplicates(S string) string {
	ans := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		n := len(ans) - 1
		if n >= 0 && ans[n] == S[i] {
			ans = ans[:n]
		} else {
			ans = append(ans, S[i])
		}
	}
	return string(ans)
}
