package problem598

func maxCount(m int, n int, ops [][]int) int {
	for _, op := range ops {
		if m > op[0] {
			m = op[0]
		}
		if n > op[1] {
			n = op[1]
		}
	}
	return m * n
}
