package problem1043

func maxSumAfterPartitioning(A []int, K int) int {
	n := len(A)
	dp := make([]int, n+1)
	// A's length grows up from 1 to n
	for l := 1; l <= n; l++ {
		m := 0 // max value of last k items in A[:l]
		for k := 1; k <= K && 0 <= l-k; k++ {
			m = max(m, A[l-k])
			sum := dp[l-k] + m*k
			dp[l] = max(dp[l], sum)
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
