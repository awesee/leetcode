package n_repeated_element_in_size_2n_array

func repeatedNTimes(A []int) int {
	m := make(map[int]bool)
	for _, v := range A {
		if m[v] {
			return v
		}
		m[v] = true
	}
	return 0
}
