package problem1021

func removeOuterParentheses(S string) string {
	ans, n := make([]rune, 0, len(S)), 0
	for _, c := range S {
		if c == '(' {
			n++
		} else {
			n--
		}
		if (c == '(' && n != 1) || (c == ')' && n != 0) {
			ans = append(ans, c)
		}
	}
	return string(ans)
}
