package problem985

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum, ans := 0, make([]int, 0)
	for _, v := range A {
		if v%2 == 0 {
			sum += v
		}
	}
	for _, q := range queries {
		v, i := q[0], q[1]
		if A[i]%2 == 0 {
			sum -= A[i]
		}
		A[i] += v
		if A[i]%2 == 0 {
			sum += A[i]
		}
		ans = append(ans, sum)
	}
	return ans
}
