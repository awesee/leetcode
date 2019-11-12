package problem867

func transpose(A [][]int) [][]int {
	m, n := len(A[0]), len(A)
	B := make([][]int, m)
	for i := 0; i < m; i++ {
		B[i] = make([]int, n)
	}
	for i, row := range A {
		for j, v := range row {
			B[j][i] = v
		}
	}
	return B
}
