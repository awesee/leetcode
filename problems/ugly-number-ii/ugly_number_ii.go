package problem264

func nthUglyNumber(n int) int {
	ans, idx := make([]int, n), [3]int{}
	ans[0] = 1
	for i := 1; i < n; i++ {
		for j, n := range [...]int{2, 3, 5} {
			if ans[idx[j]]*n <= ans[i-1] {
				idx[j]++
			}
			if num := ans[idx[j]] * n; j == 0 {
				ans[i] = num
			} else if ans[i] > num {
				ans[i] = num
			}
		}
	}
	return ans[n-1]
}
