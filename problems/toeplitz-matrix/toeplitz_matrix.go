package problem766

func isToeplitzMatrix(matrix [][]int) bool {
	for i, row := range matrix[1:] {
		for j, v := range row[1:] {
			if v != matrix[i][j] {
				return false
			}
		}
	}
	return true
}
