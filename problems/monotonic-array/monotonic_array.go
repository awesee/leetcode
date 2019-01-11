package monotonic_array

func isMonotonic(A []int) bool {
	t, f := 0, 0
	for i, v := range A[1:] {
		if v > A[i] {
			t++
		} else if v < A[i] {
			f++
		}
	}
	return t == 0 || f == 0
}
