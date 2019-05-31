package remove_all_adjacent_duplicates_in_string

func removeDuplicates(S string) string {
	ans := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		n := len(ans)
		if n >= 1 && ans[n-1] == S[i] {
			ans = ans[:n-1]
		} else {
			ans = append(ans, S[i])
		}
	}
	return string(ans)
}
