package problem922

func sortArrayByParityII(A []int) []int {
	i, l := 0, len(A)
	for j := 1; j < l; j += 2 {
		if A[j]&1 == 0 {
			for A[i]&1 == 0 {
				i += 2
			}
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
