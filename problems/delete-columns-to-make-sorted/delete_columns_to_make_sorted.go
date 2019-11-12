package problem944

func minDeletionSize(A []string) int {
	ans := 0
	for i := 0; i < len(A[0]); i++ {
		for j, str := range A[1:] {
			if str[i] < A[j][i] {
				ans++
				break
			}
		}
	}
	return ans
}
