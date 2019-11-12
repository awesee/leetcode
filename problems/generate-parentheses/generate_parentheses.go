package problem22

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	next("(", 1, 0, n, &ans)
	return ans
}

func next(s string, l, r, n int, ans *[]string) {
	if l == n && r == n {
		*ans = append(*ans, s)
	}
	if l < n {
		next(s+"(", l+1, r, n, ans)
	}
	if r < l {
		next(s+")", l, r+1, n, ans)
	}
}
