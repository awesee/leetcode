package problem661

func imageSmoother(M [][]int) [][]int {
	r, c := len(M), len(M[0])
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
		for j := 0; j < c; j++ {
			sum, count := 0, 0
			for m := i - 1; m <= i+1; m++ {
				for n := j - 1; n <= j+1; n++ {
					if m >= 0 && n >= 0 && m < r && n < c {
						sum += M[m][n]
						count++
					}
				}
			}
			ans[i][j] = sum / count
		}
	}
	return ans
}
