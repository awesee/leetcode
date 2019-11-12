package problem413

func numberOfArithmeticSlices(A []int) int {
	ans, i, j, l := 0, 0, 0, len(A)-1
	for i < l-1 {
		j = i + 1
		for j < l && A[j+1]-A[j] == A[j]-A[j-1] {
			j++
		}
		i, j = j, j-i
		ans += j * (j - 1) / 2
	}
	return ans
}
