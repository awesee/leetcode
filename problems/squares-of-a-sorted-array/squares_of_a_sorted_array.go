package squares_of_a_sorted_array

import "sort"

func sortedSquares(A []int) []int {
	ans := make([]int, len(A))
	for i, v := range A {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}
