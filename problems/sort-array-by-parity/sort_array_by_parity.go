package problem905

import "sort"

func sortArrayByParity(A []int) []int {
	sort.Slice(A, func(i, j int) bool {
		return A[i]%2 < A[j]%2
	})
	return A
}
