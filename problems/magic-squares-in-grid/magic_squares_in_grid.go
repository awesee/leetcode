package problem840

func numMagicSquaresInside(grid [][]int) int {
	ans, r, c := 0, len(grid)-1, len(grid[0])-1
	for i := 1; i < r; i++ {
	flag:
		for j := 1; j < c; j++ {
			if grid[i][j] == 5 &&
				grid[i-1][j-1]+grid[i-1][j]+grid[i-1][j+1] == 15 &&
				grid[i+1][j-1]+grid[i+1][j]+grid[i+1][j+1] == 15 &&
				grid[i-1][j-1]+grid[i][j-1]+grid[i+1][j-1] == 15 &&
				grid[i-1][j+1]+grid[i][j+1]+grid[i+1][j+1] == 15 &&
				grid[i-1][j]+grid[i+1][j] == 10 &&
				grid[i][j-1]+grid[i][j+1] == 10 &&
				grid[i-1][j-1]+grid[i+1][j+1] == 10 &&
				grid[i-1][j+1]+grid[i+1][j-1] == 10 {
				for m := i - 1; m <= i+1; m++ {
					for n := j - 1; n <= j+1; n++ {
						if (m != i && n != j && grid[m][n] == 5) || grid[m][n] < 1 || grid[m][n] > 9 {
							continue flag
						}
					}
				}
				ans++
			}
		}
	}
	return ans
}
