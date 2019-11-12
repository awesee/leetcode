package problem454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	ans, n := 0, len(A)
	sum := make(map[int]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[A[i]+B[j]]++
		}
	}
	for k := 0; k < n; k++ {
		for l := 0; l < n; l++ {
			ans += sum[-C[k]-D[l]]
		}
	}
	return ans
}
