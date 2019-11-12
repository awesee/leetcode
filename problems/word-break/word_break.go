package problem139

func wordBreak(s string, wordDict []string) bool {
	wm, l := make(map[string]bool), len(s)
	for _, w := range wordDict {
		wm[w] = true
	}
	dp := make([]bool, l+1)
	dp[0] = true
	for i := 1; i <= l; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wm[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[l]
}
