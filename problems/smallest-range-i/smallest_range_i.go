package problem908

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]
	for _, v := range A {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	ans := max - min - 2*K
	if ans < 0 {
		return 0
	}
	return ans
}
