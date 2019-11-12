package problem896

func isMonotonic(A []int) bool {
	increasing, decreasing := true, true
	for i, v := range A[1:] {
		if v > A[i] {
			decreasing = false
		} else if v < A[i] {
			increasing = false
		}
	}
	return increasing || decreasing
}
