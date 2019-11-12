package problem119

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0], row[rowIndex] = 1, 1
	for i := 1; i < rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}
