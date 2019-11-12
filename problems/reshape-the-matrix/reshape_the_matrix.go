package problem566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	m, n := len(nums), len(nums[0])
	if m*n != r*c {
		return nums
	}
	ans := make([][]int, r)
	for i := 0; i < r; i++ {
		ans[i] = make([]int, c)
	}
	i, j := 0, 0
	for _, row := range nums {
		for _, v := range row {
			ans[i][j] = v
			if j++; j == c {
				i++
				j = 0
			}
		}
	}
	return ans
}
