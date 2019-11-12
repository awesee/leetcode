package problem832

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		for i, j := 0, len(row)-1; i <= j; i, j = i+1, j-1 {
			row[i], row[j] = row[j]^1, row[i]^1
		}
	}
	return A
}
