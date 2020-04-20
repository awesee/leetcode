package problem32

func longestValidParentheses(s string) int {
	ans, l := 0, len(s)
	dp := make([]int, l)
	for i, c := range s {
		if i == 0 || c == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2
			if i > 1 {
				dp[i] += dp[i-2]
			}
		} else if i > dp[i-1] && s[i-1-dp[i-1]] == '(' {
			dp[i] = dp[i-1] + 2
			if i-1 > dp[i-1] {
				dp[i] += dp[i-2-dp[i-1]]
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
