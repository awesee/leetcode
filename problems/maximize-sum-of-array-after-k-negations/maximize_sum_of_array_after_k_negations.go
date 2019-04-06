package maximize_sum_of_array_after_k_negations

import "sort"

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)
	ans, min := 0, A[len(A)-1]
	for _, v := range A {
		if K > 0 && v <= 0 {
			K, v = K-1, -v
		}
		if v < min {
			min = v
		}
		ans += v
	}
	if K&1 == 1 {
		ans -= min * 2
	}
	return ans
}
