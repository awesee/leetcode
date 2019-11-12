package problem48

func rotate(matrix [][]int) {
	l := len(matrix)
	t := make([][]int, l)
	for i := 0; i < l; i++ {
		t[i] = make([]int, l)
		copy(t[i], matrix[i])
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			matrix[j][l-1-i] = t[i][j]
		}
	}
}
