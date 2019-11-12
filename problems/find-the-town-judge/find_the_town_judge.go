package problem997

func findJudge(N int, trust [][]int) int {
	ans, count := 1, make([]int, N+1)
	for _, t := range trust {
		count[t[0]] = N
		count[t[1]]++
		if count[t[1]] == N-1 {
			ans = t[1]
		}
		if ans == t[0] {
			ans = -1
		}
	}
	return ans
}
